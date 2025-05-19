import axios from 'axios';

const API_BASE = process.env.REACT_APP_MARKET_API_BASE_URL;
const API_TIMEOUT = process.env.REACT_APP_API_TIMEOUT;

const api = axios.create({
  baseURL: API_BASE,
  timeout: parseInt(API_TIMEOUT),
});

// Добавляем интерцептор для авторизации
api.interceptors.request.use(config => {
  const token = getAuthToken();
  if (token) {
    config.headers['api_key'] = token;
  }
  return config;
});

export const getBalance = async () => {
  try {
    const response = await api.get('/account/balance');
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const createDeposit = async (currency) => {
  try {
    const response = await api.post('/transactions/deposit', {
      currency
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const createWithdraw = async (currency, amount, address) => {
  try {
    const response = await api.post('/transactions/withdraw', {
      currency,
      amount,
      address
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const createBid = async (bidData) => {
  try {
    const response = await api.post('/bid', bidData);
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Вспомогательная функция для получения токена
const getAuthToken = () => {
  const cookies = document.cookie.split(';');
  const tokenCookie = cookies.find(c => c.trim().startsWith('token='));
  return tokenCookie ? tokenCookie.split('=')[1] : null;
};

export const getTransactionsHistory = async (filters = {}) => {
  try {
    const response = await api.get('/transactions/transfers', {
      params: filters
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const getTradesHistory = async (filters = {}) => {
  try {
    const response = await api.get('/bids', {
      params: filters
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const cancelTrade = async (tradeId) => {
  try {
    console.log(`/market/${tradeId}`);
    const response = await api.delete(`/market/${tradeId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};