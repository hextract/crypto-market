import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { isAuthenticated, logout } from '../../api/authService';
import {
  getBalance,
  createDeposit,
  createWithdraw,
  createBid,
  getMarketData
} from '../../api/marketService';
import './Main.css';
import logo from '../../assets/logo-purple.svg';

const Main = () => {
  const navigate = useNavigate();
  const [isAuth, setIsAuth] = useState(null);
  const [tradeType, setTradeType] = useState("buy");
  const [isCopied, setIsCopied] = useState(false);
  const [formData, setFormData] = useState({
    minPrice: "",
    maxPrice: "",
    amount: "",
    speed: ""
  });
  const [errors, setErrors] = useState({
    minPrice: null,
    maxPrice: null,
    amount: null,
    speed: null
  });
  const [submitAttempted, setSubmitAttempted] = useState(false);
  const [balances, setBalances] = useState({ USDT: 0, BTC: 0 });
  const [marketData, setMarketData] = useState({
    pair: 'USDT/BTC',
    price: 0,
    high: 0,
    change: 0
  });
  const [errorModal, setErrorModal] = useState({
    show: false,
    message: ''
  });

  // States for modals
  const [showWithdrawModal, setShowWithdrawModal] = useState(false);
  const [showDepositModal, setShowDepositModal] = useState(false);
  const [withdrawCurrency, setWithdrawCurrency] = useState('USDT');
  const [withdrawForm, setWithdrawForm] = useState({
    amount: '',
    wallet: ''
  });
  const [withdrawErrors, setWithdrawErrors] = useState({
    amount: null,
    wallet: null
  });
  const [depositAddress, setDepositAddress] = useState('');

  useEffect(() => {
    const checkAuth = async () => {
      const auth = await isAuthenticated();
      setIsAuth(auth);
      if (!auth) navigate('/login');
    };
    checkAuth();

    // Загружаем баланс и данные рынка
    if (isAuth) {
      loadData();
    }
  }, [navigate, isAuth]);

  const loadData = async () => {
    try {
      // Загружаем баланс
      const balanceResponse = await getBalance();
      const newBalances = { USDT: 0, BTC: 0 };
      balanceResponse.forEach(item => {
        newBalances[item.currency] = item.amount + 100;
      });
      setBalances(newBalances);

      // Загружаем данные рынка
      const marketResponse = await getMarketData();
      setMarketData({
        pair: 'USDT/BTC',
        price: marketResponse.current_price,
        high: marketResponse.high_24h,
        change: marketResponse.price_change_percentage_24h
      });
    } catch (error) {
      console.error('Failed to load data:', error);
      showError('Failed to load data. Please try again later.');
    }
  };

  const showError = (message) => {
    setErrorModal({
      show: true,
      message
    });
  };

  const closeErrorModal = () => {
    setErrorModal({
      show: false,
      message: ''
    });
  };

  const copyToClipboard = () => {
    navigator.clipboard.writeText(depositAddress)
      .then(() => {
        setIsCopied(true);
        setTimeout(() => setIsCopied(false), 1500);
      })
      .catch(err => {
        console.error('Failed to copy: ', err);
        showError('Failed to copy address');
      });
  };

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  const handleTradeTypeChange = (type) => {
    setTradeType(type);
  };

  const validateField = (name, value) => {
    if (name === "speed") {
      const numValue = parseFloat(value);
      if (!value.trim()) return "Enter stocks/hour";
      if (isNaN(numValue)) return "Must be a number";
      if (numValue <= 0) return "Must be positive";
      return null;
    }

    const numValue = parseFloat(value);
    if (!value.trim()) return "This field is required";
    if (isNaN(numValue)) return "Must be a number";
    if (numValue <= 0) return "Must be positive";
    return null;
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));

    if (submitAttempted) {
      setErrors(prev => ({ ...prev, [name]: validateField(name, value) }));
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setSubmitAttempted(true);

    const newErrors = {
      minPrice: validateField("minPrice", formData.minPrice),
      maxPrice: validateField("maxPrice", formData.maxPrice),
      amount: validateField("amount", formData.amount),
      speed: validateField("speed", formData.speed)
    };

    setErrors(newErrors);

    if (Object.values(newErrors).some(error => error)) {
      return;
    }

    try {
      const bidData = {
        from_currency: tradeType === 'buy' ? 'USDT' : 'BTC',
        to_currency: tradeType === 'buy' ? 'BTC' : 'USDT',
        min_price: parseFloat(formData.minPrice),
        max_price: parseFloat(formData.maxPrice),
        amount_to_buy: parseFloat(formData.amount),
        buy_speed: parseFloat(formData.speed)
      };

      await createBid(bidData);
      // Обновляем баланс после успешной сделки
      await loadData();
      // Можно показать уведомление об успехе
      showError('Order created successfully!');
    } catch (error) {
      console.error('Failed to create bid:', error);
      let errorMessage = 'Failed to create order';
      if (error.response) {
        if (error.response.status === 400) {
          errorMessage = error.response.data.error_message || 'Invalid data';
        } else if (error.response.status === 403) {
          errorMessage = 'Insufficient balance';
        }
      }
      showError(errorMessage);
    }
  };

  // Modal handlers
  const handleWithdrawClick = (currency) => {
    setWithdrawCurrency(currency);
    setShowWithdrawModal(true);
  };

  const handleDepositClick = async (currency) => {
    try {
      const response = await createDeposit(currency);
      setDepositAddress(response.address);
      setShowDepositModal(true);
    } catch (error) {
      console.error('Failed to get deposit address:', error);
      showError('Failed to get deposit address. Please try again.');
    }
  };

  const closeModal = () => {
    setShowWithdrawModal(false);
    setShowDepositModal(false);
    setWithdrawForm({ amount: '', wallet: '' });
    setWithdrawErrors({ amount: null, wallet: null });
    setIsCopied(false);
  };

  const handleWithdrawChange = (e) => {
    const { name, value } = e.target;
    setWithdrawForm(prev => ({ ...prev, [name]: value }));

    if (name === 'amount') {
      const numValue = parseFloat(value);
      if (!value.trim()) {
        setWithdrawErrors(prev => ({ ...prev, amount: 'Amount is required' }));
      } else if (isNaN(numValue)) {
        setWithdrawErrors(prev => ({ ...prev, amount: 'Must be a number' }));
      } else if (numValue <= 0) {
        setWithdrawErrors(prev => ({ ...prev, amount: 'Must be positive' }));
      } else if (numValue > balances[withdrawCurrency]) {
        setWithdrawErrors(prev => ({ ...prev, amount: 'Insufficient balance' }));
      } else {
        setWithdrawErrors(prev => ({ ...prev, amount: null }));
      }
    } else if (name === 'wallet' && !value.trim()) {
      setWithdrawErrors(prev => ({ ...prev, wallet: 'Wallet is required' }));
    } else if (name === 'wallet') {
      setWithdrawErrors(prev => ({ ...prev, wallet: null }));
    }
  };

  const handleWithdrawSubmit = async (e) => {
    e.preventDefault();

    const newErrors = {
      amount: !withdrawForm.amount ? 'Amount is required' :
        isNaN(parseFloat(withdrawForm.amount)) ? 'Must be a number' :
          parseFloat(withdrawForm.amount) <= 0 ? 'Must be positive' :
            parseFloat(withdrawForm.amount) > balances[withdrawCurrency] ? 'Insufficient balance' : null,
      wallet: !withdrawForm.wallet.trim() ? 'Wallet is required' : null
    };

    setWithdrawErrors(newErrors);

    if (newErrors.amount || newErrors.wallet) {
      return;
    }

    try {
      await createWithdraw(
        withdrawCurrency,
        parseFloat(withdrawForm.amount),
        withdrawForm.wallet
      );
      // Обновляем баланс после успешного вывода
      await loadData();
      closeModal();
      showError('Withdrawal request created successfully!');
    } catch (error) {
      console.error('Failed to create withdrawal:', error);
      let errorMessage = 'Failed to create withdrawal';
      if (error.response) {
        if (error.response.status === 400) {
          errorMessage = error.response.data.error_message || 'Invalid data';
        } else if (error.response.status === 403) {
          errorMessage = 'Insufficient balance';
        }
      }
      showError(errorMessage);
    }
  };

  if (isAuth === null) {
    return <div className="loading-screen">Loading...</div>;
  }

  return (
    <div className="main-container">
      <nav className="navbar-main">
        <div className="logo">
          <img src={logo} alt="Logo" className="logo-purple"/>
          <span>CONT</span>
        </div>
        <div className="nav-links">
          <a href="/main" className="active">trade</a>
          <a href="/profile">profile</a>
          <button onClick={handleLogout} className="logout-btn">Logout</button>
        </div>
      </nav>

      <div className="content">
        <div className="chart-section">
          <p className="chart-placeholder">тут график</p>
        </div>

        <div className="trade-panel">
          <div className="market-info">
            <div className="main-info">
              <span className="trading-pair">{marketData.pair}</span>
              <span className="price">{marketData.price.toFixed(6)}</span>
            </div>
            <div className="add-info">
              <p className="high-price">high this day: {marketData.high.toFixed(6)}</p>
              <p className={`price-change ${marketData.change >= 0 ? 'positive' : 'negative'}`}>
                {marketData.change >= 0 ? '+' : ''}{marketData.change.toFixed(2)}% this day
              </p>
            </div>
          </div>

          <div className="market-data">
            <div className="button-group">
              <button
                className={tradeType === "buy" ? "active" : ""}
                onClick={() => handleTradeTypeChange("buy")}
              >
                buy
              </button>
              <button
                className={tradeType === "sell" ? "active" : ""}
                onClick={() => handleTradeTypeChange("sell")}
              >
                sell
              </button>
            </div>

            <form className="main-form" onSubmit={handleSubmit}>
              <div className="input-container-main">
                <input
                  type="text"
                  name="minPrice"
                  placeholder="min price"
                  value={formData.minPrice}
                  onChange={handleChange}
                  className={`input ${errors.minPrice ? "error" : ""}`}
                />
                {errors.minPrice && submitAttempted && (
                  <div className="tooltip show">{errors.minPrice}</div>
                )}
              </div>

              <div className="input-container-main">
                <input
                  type="text"
                  name="maxPrice"
                  placeholder="max price"
                  value={formData.maxPrice}
                  onChange={handleChange}
                  className={`input ${errors.maxPrice ? "error" : ""}`}
                />
                {errors.maxPrice && submitAttempted && (
                  <div className="tooltip show">{errors.maxPrice}</div>
                )}
              </div>

              <div className="input-container-main">
                <input
                  type="text"
                  name="amount"
                  placeholder={`amount to ${tradeType}`}
                  value={formData.amount}
                  onChange={handleChange}
                  className={`input ${errors.amount ? "error" : ""}`}
                />
                {errors.amount && submitAttempted && (
                  <div className="tooltip show">{errors.amount}</div>
                )}
              </div>

              <div className="input-container-main">
                <input
                  type="text"
                  name="speed"
                  placeholder="buy speed (stocks/hour)"
                  value={formData.speed}
                  onChange={handleChange}
                  className={`input ${errors.speed ? "error" : ""}`}
                />
                {errors.speed && submitAttempted && (
                  <div className="tooltip show">{errors.speed}</div>
                )}
              </div>

              <button
                type="submit"
                className={`proceed-btn ${Object.values(errors).some(e => e) ? "disabled" : "active"}`}
                disabled={Object.values(errors).some(e => e)}
              >
                proceed
              </button>
            </form>
          </div>
        </div>
      </div>

      <div className="balance-section">
        <div className="balance-card">
          <span className="balance-currency">USDT: {balances.USDT.toFixed(2)}</span>
          <div className="balance-button">
            <button
              className="balance-btn withdraw-btn"
              onClick={() => handleWithdrawClick('USDT')}
            >
              withdraw
            </button>
            <button
              className="balance-btn deposit-btn"
              onClick={() => handleDepositClick('USDT') }
            >
              deposit
            </button>
          </div>
        </div>

        <div className="balance-card">
          <span className="balance-currency">BTC: {balances.BTC.toFixed(6)}</span>
          <div className="balance-button">
            <button
              disabled={true}
              className="balance-btn withdraw-btn disabled"
              onClick={() => handleWithdrawClick('BTC')}
            >
              withdraw
            </button>
            <button
              disabled={true}
              className="balance-btn deposit-btn disabled"
              onClick={() => handleDepositClick('BTC') }
            >
              deposit
            </button>
          </div>
        </div>
      </div>

      {/* Withdraw Modal */}
      {showWithdrawModal && (
        <div className="modal-overlay">
          <div className="modal">
            <div className="modal-header">
              <h3>Withdraw {withdrawCurrency}</h3>
              <button className="modal-close" onClick={closeModal}>×</button>
            </div>
            <form onSubmit={handleWithdrawSubmit} className="modal-form">
              <div className="modal-input-container">
                <input
                  type="text"
                  name="amount"
                  placeholder={`Amount (${withdrawCurrency})`}
                  value={withdrawForm.amount}
                  onChange={handleWithdrawChange}
                  className={`modal-input ${withdrawErrors.amount ? 'error' : ''}`}
                />
                {withdrawErrors.amount && (
                  <div className="modal-tooltip">{withdrawErrors.amount}</div>
                )}
              </div>

              <div className="modal-input-container">
                <input
                  type="text"
                  name="wallet"
                  placeholder="Your wallet address"
                  value={withdrawForm.wallet}
                  onChange={handleWithdrawChange}
                  className={`modal-input ${withdrawErrors.wallet ? 'error' : ''}`}
                />
                {withdrawErrors.wallet && (
                  <div className="modal-tooltip">{withdrawErrors.wallet}</div>
                )}
              </div>

              <button
                type="submit"
                className="modal-submit-btn"
                disabled={!!withdrawErrors.amount || !!withdrawErrors.wallet}
              >
                Confirm Withdraw
              </button>
            </form>
          </div>
        </div>
      )}

      {/* Deposit Modal */}
      {showDepositModal && (
        <div className="modal-overlay">
          <div className="modal">
            <div className="modal-header">
              <h3>Deposit</h3>
              <button className="modal-close" onClick={closeModal}>×</button>
            </div>
            <div className="modal-content">
              <p>Send funds to the following address:</p>
              <div className="wallet-address-container">
                <div className="wallet-address">
                  {depositAddress}
                </div>
                <button
                  className="copy-btn"
                  onClick={copyToClipboard}
                  title="Copy to clipboard"
                >
                  {isCopied ? 'Copied!' : 'Copy'}
                </button>
              </div>
              <p className="wallet-note">
                Please make sure you're sending the correct currency to this address.
                The deposit will be credited after 3 network confirmations.
              </p>
            </div>
          </div>
        </div>
      )}

      {/* Error Modal */}
      {errorModal.show && (
        <div className="modal-overlay" onClick={closeErrorModal}>
          <div className="modal" onClick={e => e.stopPropagation()}>
            <div className="modal-header">
              <h3>{errorModal.message.includes('success') ? 'Success' : 'Error'}</h3>
              <button className="modal-close" onClick={closeErrorModal}>×</button>
            </div>
            <div className="modal-content">
              <p>{errorModal.message}</p>
              <div className="modal-buttons">
                <button
                  onClick={closeErrorModal}
                  className="modal-confirm-btn"
                >
                  OK
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default Main;