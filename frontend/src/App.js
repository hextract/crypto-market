import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import PublicRoute from './components/PublicRoute';
import AuthRoute from './components/AuthRoute';
import About from './pages/About/About';
import Login from './pages/Login/Login';
import Register from './pages/Register/Register';
import Main from './pages/Main/Main';
import Profile from "./pages/Profile/Profile";
import LanguageSwitcher from './components/LanguageSwitcher';
import './App.css';

function App() {
  return (
    <Router>
      <LanguageSwitcher />
      <Routes>
        <Route path="/" element={<PublicRoute><About /></PublicRoute>} />
        <Route path="/about" element={<PublicRoute><About /></PublicRoute>} />
        <Route path="/login" element={<PublicRoute><Login /></PublicRoute>} />
        <Route path="/register" element={<PublicRoute><Register /></PublicRoute>} />
        <Route path="/main/*" element={<AuthRoute><Main /></AuthRoute>} />
        <Route path="/profile/*" element={<AuthRoute><Profile /></AuthRoute>} />
        <Route path="*" element={<AuthRoute><Main /></AuthRoute>} />
      </Routes>
    </Router>
  );
}

export default App;