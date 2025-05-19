import "./About.css";
import { useNavigate } from "react-router-dom";
import { isAuthenticated } from '../../api/authService';
import { useEffect } from 'react';
import logo from "../../assets/logo.svg";
import { useTranslation } from 'react-i18next';

const About = () => {
  const navigate = useNavigate();
  const { t } = useTranslation();

  useEffect(() => {
    const checkAuthAndRedirect = async () => {
      const auth = await isAuthenticated();
      if (auth) navigate('/main');
    };
    checkAuthAndRedirect();
  }, [navigate]);

  return (
    <div className="about-page">
      <div className="about-wrapper">
        <div className="about-content">
          <div className="logo-container">
            <img src={logo} alt="Logo" className="logo" />
          </div>
          <div className="text-container">
            <h1 className="title-about">{t('about.title')}</h1>
            <p className="text-about">
              {t('about.description')}
            </p>
          </div>
          <div className="button-container">
            <button className="btn register" onClick={() => navigate("/register")}>{t('about.register')}</button>
            <button className="btn signin" onClick={() => navigate("/login")}>{t('about.signIn')}</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default About;