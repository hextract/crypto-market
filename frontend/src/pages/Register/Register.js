import React, {useState, useEffect, useCallback} from 'react';
import { useNavigate } from 'react-router-dom';
import { registerUser, setAuthToken } from '../../api/authService';
import './Register.css';
import logo from '../../assets/logo-purple.svg';

const Register = () => {
    const [formData, setFormData] = useState({
        email: '',
        password: ''
    });
    const [errors, setErrors] = useState({
        email: null,
        password: null,
        server: null
    });
    const [isLoading, setIsLoading] = useState(false);
    const [submitAttempted, setSubmitAttempted] = useState(false);
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

    const validateForm = useCallback(() => {
        const newErrors = {
            email: !formData.email ? 'Email is required' :
              !validateEmail(formData.email) ? 'Invalid email format' : null,
            password: !formData.password ? 'Password is required' :
              formData.password.length < 8 ? 'Min 8 characters required' : null,
            server: null
        };
        setErrors(newErrors);
        return !newErrors.email && !newErrors.password;
    }, [formData]);

    useEffect(() => {
        if (submitAttempted) {
            validateForm();
        }
    }, [submitAttempted, validateForm]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        e.stopPropagation();

        if (!validateForm()) {
            return;
        }
        setIsLoading(true);

        try {
            const data = await registerUser(formData.email, formData.email, formData.password);
            setAuthToken(data.token);
            navigate('/main');
        } catch (err) {
            if (err.response) {
                if (err.response.status === 409) {
                    setErrors(prev => ({ ...prev, server: 'This email is already registered' }));
                } else if (err.response.status >= 500) {
                    setErrors(prev => ({ ...prev, server: 'Server is unavailable, please try later' }));
                }
            } else {
                setErrors(prev => ({ ...prev, server: 'Network error' }));
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

    return (
      <div className="container">
          <nav className="navbar">
              <div className="logo">
                  <img src={logo} alt="Logo" className="logo-purple" />
                  <span className="title-navbar">CONT</span>
              </div>
              <div className="nav-links">
                  <a className="link" href="/about">about</a>
                  <a className="link" href="/login">login</a>
                  <a className="link active" href="/register">register</a>
              </div>
          </nav>

          <h2 className="title">Sign up</h2>

          {errors.server && <div className="error-message">{errors.server}</div>}

          <form onSubmit={handleSubmit} className="input-wrapper">
              <div className="input-container">
                  <input
                    formNoValidate={true}
                    type="email"
                    name="email"
                    placeholder="email"
                    className={`input-register ${errors.email ? 'error' : ''}`}
                    value={formData.email}
                    onChange={handleChange}
                    autoComplete="username"
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
                    placeholder="password"
                    className={`input-register ${errors.password ? 'error' : ''}`}
                    value={formData.password}
                    onChange={handleChange}
                    minLength="8"
                    autoComplete="new-password"
                  />
                  {errors.password && submitAttempted && (
                    <div className="tooltip show">{errors.password}</div>
                  )}
              </div>

              <button
                type="submit"
                className={`sign-button ${!isFormValid() || isLoading ? 'disabled' : 'active'}`}
                disabled={!isFormValid() || isLoading}
              >
                  {isLoading ? 'Loading...' : 'Sign up!'}
              </button>
          </form>
      </div>
    );
};

export default Register;