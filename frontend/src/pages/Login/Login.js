import React, {useCallback, useEffect, useState} from "react";
import { useNavigate } from "react-router-dom";
import { loginUser, setAuthToken } from "../../api/authService";
import "../Register/Register.css";
import logo from "../../assets/logo-purple.svg";
import eyeOpen from '../../assets/eye-open.png';
import eyeClosed from '../../assets/eye-closed.png';
import { useTranslation } from 'react-i18next';

const LoginPage = () => {
    const [formData, setFormData] = useState({
        email: '',
        password: ''
    });

    const [errors, setErrors] = useState({
        email: null,
        password: null,
        server: null
    });

    const [submitAttempted, setSubmitAttempted] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const [showPassword, setShowPassword] = useState(false);
    const navigate = useNavigate();
    const { t } = useTranslation();

    const validateEmail = (email) => {
        const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return re.test(email);
    };

    const isFormValid = () => {
        return (
          formData.email &&
          formData.password &&
          validateEmail(formData.email) &&
          formData.password.length >= 8
        );
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        e.stopPropagation();

        if (!validateForm()) {
            return;
        }

        setIsLoading(true);

        try {
            const data = await loginUser(formData.email, formData.password);
            setAuthToken(data.token);
            navigate('/main');
        } catch (err) {
            if (err.response) {
                if (err.response.status === 401) {
                    setErrors(prev => ({ ...prev, server: t('auth.errors.incorrectCredentials') }));
                } else if (err.response.status >= 500) {
                    setErrors(prev => ({ ...prev, server: t('auth.errors.serverUnavailable') }));
                }
            } else {
                setErrors(prev => ({ ...prev, server: t('auth.errors.networkError') }));
            }
        } finally {
            setIsLoading(false);
        }
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData(prev => ({ ...prev, [name]: value }));

        if ((formData.email && name !== "email") || (formData.password && name !== "password")) {
            setSubmitAttempted(true);
        }
        if (submitAttempted) {
            validateForm();
        }
    };

    const validateForm = useCallback(() => {
        const newErrors = {
            email: !formData.email ? t('auth.errors.emailRequired') :
              !validateEmail(formData.email) ? t('auth.errors.invalidEmail') : null,
            password: !formData.password ? t('auth.errors.passwordRequired') :
              formData.password.length < 8 ? t('auth.errors.minPassword') : null,
            server: null,
        };
        setErrors(newErrors);
        return !newErrors.email && !newErrors.password;
    }, [formData, t]);

    useEffect(() => {
        if (submitAttempted) {
            validateForm();
        }
    }, [submitAttempted, validateForm]);

    const toggleShowPassword = () => {
        setShowPassword(!showPassword);
    };

    return (
      <div className="container">
          <nav className="navbar">
              <div className="logo">
                  <img src={logo} alt="Logo" className="logo-purple"/>
                  <span className="title-navbar">{t('navbar.logo')}</span>
              </div>
              <div className="nav-links">
                  <a className="link" href="/about">{t('navbar.about')}</a>
                  <a className="link active" href="/login">{t('navbar.login')}</a>
                  <a className="link" href="/register">{t('navbar.register')}</a>
              </div>
          </nav>

          <h2 className="title">{t('auth.signIn')}</h2>

          {errors.server && <div className="error-message">{errors.server}</div>}

          <form onSubmit={handleSubmit} className="input-wrapper">
              <div className="input-container">
                  <input
                    formNoValidate={true}
                    type="email"
                    name="email"
                    placeholder={t('auth.email')}
                    className={`input-register ${errors.email ? "error" : ""}`}
                    value={formData.email}
                    onChange={handleChange}
                  />
                  {errors.email && submitAttempted && (
                    <div className="tooltip show">{errors.email}</div>
                  )}
              </div>

              <div className="input-container password-container">
                  <input
                    formNoValidate={true}
                    type={showPassword ? "text" : "password"}
                    name="password"
                    minLength="8"
                    placeholder={t('auth.password')}
                    className={`input-register ${errors.password ? "error" : ""}`}
                    value={formData.password}
                    onChange={handleChange}
                  />
                  <button
                    type="button"
                    className="password-toggle"
                    onClick={toggleShowPassword}
                    aria-label={showPassword ? t('auth.hidePassword') : t('auth.showPassword')}
                  >
                      <img
                        src={showPassword ? eyeOpen : eyeClosed}
                        alt={showPassword ? t('auth.hidePassword') : t('auth.showPassword')}
                        className="eye-icon"
                      />
                  </button>
                  {errors.password && submitAttempted && (
                    <div className="tooltip show">{errors.password}</div>
                  )}
              </div>

              <button
                type="submit"
                className={`sign-button ${!isFormValid() || isLoading ? "disabled" : "active"}`}
                disabled={!isFormValid() || isLoading}
              >
                  {isLoading ? t('auth.loading') : t('auth.signInBtn')}
              </button>
          </form>
      </div>
    );
};

export default LoginPage;