import { useEffect, useState } from 'react';
import { Navigate } from 'react-router-dom';
import { isAuthenticated } from '../api/authService';
import { useTranslation } from 'react-i18next';

const AuthRoute = ({ children }) => {
  const [authStatus, setAuthStatus] = useState(null);
  const { t } = useTranslation();

  useEffect(() => {
    const checkAuth = async () => {
      const isAuth = await isAuthenticated();
      setAuthStatus(isAuth);
    };
    checkAuth();
  }, []);

  if (authStatus === null) {
    return <div>{t('auth.loading')}</div>;
  }

  return authStatus ? children : <Navigate to="/login" replace />;
};

export default AuthRoute;