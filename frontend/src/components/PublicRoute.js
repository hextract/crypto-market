import { useEffect, useState } from 'react';
import { Navigate } from 'react-router-dom';
import { isAuthenticated } from '../api/authService';

const PublicRoute = ({ children }) => {
  const [authStatus, setAuthStatus] = useState(null);

  useEffect(() => {
    const checkAuth = async () => {
      const isAuth = await isAuthenticated();
      setAuthStatus(isAuth);
    };
    checkAuth();
  }, []);

  if (authStatus === null) {
    return <div>Loading...</div>; // или ваш лоадер
  }

  return !authStatus ? children : <Navigate to="/main" replace />;
};

export default PublicRoute;