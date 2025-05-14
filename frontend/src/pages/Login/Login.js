import React, {useCallback, useEffect, useState} from "react";
import { useNavigate } from "react-router-dom";
import { loginUser, setAuthToken } from "../../api/authService";
import "../Register/Register.css";
import logo from "../../assets/logo-purple.svg";

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
    const navigate = useNavigate();

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
                    setErrors(prev => ({ ...prev, server: "Incorrect email or password" }));
                } else if (err.response.status >= 500) {
                    setErrors(prev => ({ ...prev, server: "Server is unavailable, please try later" }));
                }
            } else {
                setErrors(prev => ({ ...prev, server: "Network error, please try later" }));
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
            email: !formData.email ? 'Email is required' :
              !validateEmail(formData.email) ? 'Invalid email format' : null,
            password: !formData.password ? 'Password is required' :
              formData.password.length < 8 ? 'Min 8 characters required' : null,
            server: null,
        };
        setErrors(newErrors);
        return !newErrors.email && !newErrors.password;
    }, [formData]);

    useEffect(() => {
        if (submitAttempted) {
            validateForm();
        }
    }, [submitAttempted, validateForm]);

    return (
      <div className="container">
          <nav className="navbar">
              <div className="logo">
                  <img src={logo} alt="Logo" className="logo-purple"/>
                  <span className="title-navbar">CONT</span>
              </div>
              <div className="nav-links">
                  <a className="link" href="/about">about</a>
                  <a className="link active" href="/login">login</a>
                  <a className="link" href="/register">register</a>
              </div>
          </nav>

          <h2 className="title">Sign in</h2>

          {errors.server && <div className="error-message">{errors.server}</div>}

          <form onSubmit={handleSubmit} className="input-wrapper">
              <div className="input-container">
                  <input
                    formNoValidate={true}
                    type="email"
                    name="email"
                    placeholder="email"
                    className={`input-register ${errors.email ? "error" : ""}`}
                    value={formData.email}
                    onChange={handleChange}
                  />
                  {errors.email && submitAttempted && (
                    <div className="tooltip show">{errors.email}</div>
                  )}
              </div>

              <div className="input-container">
                  <input
                    formNoValidate={true}
                    type="password"
                    name="password"
                    minLength="8"
                    placeholder="password"
                    className={`input-register ${errors.password ? "error" : ""}`}
                    value={formData.password}
                    onChange={handleChange}
                  />
                  {errors.password && submitAttempted && (
                    <div className="tooltip show">{errors.password}</div>
                  )}
              </div>

              <button
                type="submit"
                className={`sign-button ${!isFormValid() || isLoading ? "disabled" : "active"}`}
                disabled={!isFormValid() || isLoading}
              >
                  {isLoading ? "Loading..." : "Sign in!"}
              </button>
          </form>
      </div>
    );
};

export default LoginPage;