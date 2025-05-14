import "./About.css";
import { useNavigate } from "react-router-dom";
import { isAuthenticated } from '../../api/authService';
import { useEffect } from 'react';
import logo from "../../assets/logo.svg";

const About = () => {
  const navigate = useNavigate();

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
            <h1 className="title-about">CONT - first continuous order trading platform</h1>
            <p className="text-about">
              Experience seamless, high-speed trading with our continuous order book. 
              Deep liquidity, ultra-low latency, and advanced security—everything you need to trade smarter. 
              Join now and stay ahead in the crypto market.
            </p>
          </div>
          <div className="button-container">
            <button className="btn register" onClick={() => navigate("/register")}>register</button>
            <button className="btn signin" onClick={() => navigate("/login")}>sign in</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default About;