import axios from 'axios';

const API_BASE = process.env.REACT_APP_API_BASE_URL;
const API_TIMEOUT = process.env.REACT_APP_API_TIMEOUT;

const api = axios.create({
  baseURL: API_BASE,
  timeout: parseInt(API_TIMEOUT),
});

export const registerUser = async (email, login, password) => {
  try {
    const response = await api.post('/auth/register', {
      password,
      login,
      email
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const loginUser = async (login, password) => {
  try {
    const response = await api.post('/auth/login', {
      password,
      login
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const verifyToken = async () => {
  try {
    const token = getAuthToken();
    if (!token) return false;

    const response = await api.post('/auth/validate-token', {
      token,
    });
    console.log(response.data["is-valid"]);
    return response.data["is-valid"];

  } catch (error) {
    return false;
  }
};

export const setAuthToken = (token) => {
  const expires = new Date();
  expires.setTime(expires.getTime() + 24 * 60 * 60 * 1000 * 5 ); // 5 дней
  document.cookie = `token=${token}; path=/; expires=${expires.toUTCString()}; Secure; SameSite=Strict`;
};

export const getAuthToken = () => {
  const cookies = document.cookie.split(';');
  const tokenCookie = cookies.find(c => c.trim().startsWith('token='));
  return tokenCookie ? tokenCookie.split('=')[1] : null;
};

export const logout = () => {
  document.cookie = 'token=; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT';
  window.location.href = '/login';
};

export const isAuthenticated = async () => {
  return await verifyToken();
};