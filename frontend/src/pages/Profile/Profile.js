import React, { useState, useEffect, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';
import { logout } from '../../api/authService';
import {
  getTransactionsHistory,
  getTradesHistory,
  cancelTrade
} from '../../api/marketService';
import './Profile.css';
import logo from '../../assets/logo-purple.svg';

const Profile = () => {
  const navigate = useNavigate();
  const [activeTab, setActiveTab] = useState('trades');
  const [expandedId, setExpandedId] = useState(null);
  const [sortConfig, setSortConfig] = useState({ key: 'date', direction: 'desc' });
  const [filters, setFilters] = useState({
    time: 'all',
    type: 'all',
    status: 'all'
  });
  const [cancelModal, setCancelModal] = useState({
    show: false,
    id: null,
    type: null
  });
  const [errorModal, setErrorModal] = useState({
    show: false,
    message: ''
  });
  const [isLoading, setIsLoading] = useState(false);
  const [transactions, setTransactions] = useState([]);
  const [trades, setTrades] = useState([]);

  const loadData = useCallback(async () => {
    setIsLoading(true);
    try {
      if (activeTab === 'transactions') {
        const params = buildTransactionFilters();
        const data = await getTransactionsHistory(params);
        setTransactions(data.map(mapTransaction));
      } else {
        const params = buildTradeFilters();
        const data = await getTradesHistory(params);
        setTrades(data.map(mapTrade));
      }
    } catch (error) {
      console.error('Failed to load data:', error);
      showError('Failed to load history. Please try again later.');
    } finally {
      setIsLoading(false);
    }
  }, [activeTab, filters]);

  useEffect(() => {
    loadData();
  }, [loadData]);

  const buildTransactionFilters = () => {
    const params = {};
    if (filters.time !== 'all') {
      const now = Math.floor(Date.now() / 1000);
      switch (filters.time) {
        case 'hour': params.date_from = now - 3600; break;
        case 'day': params.date_from = now - 86400; break;
        case 'week': params.date_from = now - 604800; break;
        case 'month': params.date_from = now - 2592000; break;
        case 'year': params.date_from = now - 31536000; break;
      }
    }
    if (filters.type !== 'all') {
      params.operation = filters.type;
    }
    if (filters.status !== 'all') {
      params.status = filters.status;
    }
    return params;
  };

  const buildTradeFilters = () => {
    const params = {};
    if (filters.time !== 'all') {
      const now = Math.floor(Date.now() / 1000);
      switch (filters.time) {
        case 'hour': params.date_from = now - 3600; break;
        case 'day': params.date_from = now - 86400; break;
        case 'week': params.date_from = now - 604800; break;
        case 'month': params.date_from = now - 2592000; break;
        case 'year': params.date_from = now - 31536000; break;
      }
    }
    if (filters.status !== 'all') {
      params.status = filters.status;
    }
    return params;
  };

  const mapTransaction = (item) => ({
    id: item.id,
    date: new Date(item.date * 1000),
    type: item.operation,
    currency: item.currency,
    amount: item.amount,
    fee: item.commission || 0,
    status: item.status,
    wallet: item.address,
    rawStatus: item.status
  });

  const mapTrade = (item) => ({
    id: item.id,
    date: new Date(item.date * 1000),
    sellCurrency: item.currency_from,
    buyCurrency: item.currency_to,
    amount: item.amount_to,
    fee: item.commission || 0,
    status: mapTradeStatus(item.status),
    completedAmount: item.amount_to || 0,
    rawStatus: item.status
  });

  const mapTradeStatus = (status) => {
    switch (status) {
      case 'finished': return 'completed';
      case 'processing': return 'partial';
      case 'pending': return 'pending';
      case 'cancelled': return 'failed';
      default: return status;
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

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  const handleSort = (key) => {
    let direction = 'desc';
    if (sortConfig.key === key && sortConfig.direction === 'desc') {
      direction = 'asc';
    }
    setSortConfig({ key, direction });
  };

  const handleFilterChange = (filterType, value) => {
    setFilters({
      ...filters,
      [filterType]: value
    });
  };

  const resetFilters = () => {
    setFilters({
      time: 'all',
      type: 'all',
      status: 'all'
    });
  };

  const sortedData = [...(activeTab === 'transactions' ? transactions : trades)].sort((a, b) => {
    if (a[sortConfig.key] < b[sortConfig.key]) {
      return sortConfig.direction === 'asc' ? -1 : 1;
    }
    if (a[sortConfig.key] > b[sortConfig.key]) {
      return sortConfig.direction === 'asc' ? 1 : -1;
    }
    return 0;
  });

  const filteredData = sortedData.filter(item => {
    const now = new Date();
    const itemDate = new Date(item.date);
    const diffHours = (now - itemDate) / (1000 * 60 * 60);

    if (filters.time === 'hour' && diffHours > 1) return false;
    if (filters.time === 'day' && diffHours > 24) return false;
    if (filters.time === 'week' && diffHours > 168) return false;
    if (filters.time === 'month' && diffHours > 720) return false;
    if (filters.time === 'year' && diffHours > 8760) return false;

    if (activeTab === 'transactions' && filters.type !== 'all' && item.type !== filters.type) return false;

    if (filters.status !== 'all' && item.rawStatus !== filters.status) return false;

    return true;
  });

  const formatDate = (date) => {
    return date.toLocaleString();
  };

  const handleCancelClick = (id) => {
    setCancelModal({
      show: true,
      id,
      type: activeTab
    });
  };

  const confirmCancel = async () => {
    try {
      await cancelTrade(cancelModal.id);
      // Обновляем данные после отмены
      await loadData();
      setCancelModal({ show: false, id: null, type: null });
    } catch (error) {
      console.error('Failed to cancel trade:', error);
      showError('Failed to cancel trade. Please try again.');
      setCancelModal({ show: false, id: null, type: null });
    }
  };

  const closeModal = () => {
    setCancelModal({ show: false, id: null, type: null });
  };

  return (
    <div className="profile-container">
      <nav className="navbar-main">
        <div className="logo">
          <img src={logo} alt="Logo" className="logo-purple"/>
          <span>CONT</span>
        </div>
        <div className="nav-links">
          <a href="/main">trade</a>
          <a href="/profile" className="active">profile</a>
          <button onClick={handleLogout} className="logout-btn">Logout</button>
        </div>
      </nav>

      <div className="profile-content">
        <div className="profile-tabs">
          <button
            className={activeTab === 'trades' ? 'active' : ''}
            onClick={() => setActiveTab('trades')}
          >
            Trades
          </button>
          <button
            className={activeTab === 'transactions' ? 'active' : ''}
            onClick={() => setActiveTab('transactions')}
          >
            Transactions
          </button>
        </div>

        <div className="filters">
          <div className="filter-group">
            <label>Time:</label>
            <select
              value={filters.time}
              onChange={(e) => handleFilterChange('time', e.target.value)}
            >
              <option value="all">All time</option>
              <option value="hour">Last hour</option>
              <option value="day">Last 24 hours</option>
              <option value="week">Last week</option>
              <option value="month">Last month</option>
              <option value="year">Last year</option>
            </select>
          </div>

          {activeTab === 'transactions' && (
            <div className="filter-group">
              <label>Type:</label>
              <select
                value={filters.type}
                onChange={(e) => handleFilterChange('type', e.target.value)}
              >
                <option value="all">All types</option>
                <option value="deposit">Deposits</option>
                <option value="withdraw">Withdrawals</option>
              </select>
            </div>
          )}

          <div className="filter-group">
            <label>Status:</label>
            <select
              value={filters.status}
              onChange={(e) => handleFilterChange('status', e.target.value)}
            >
              <option value="all">All statuses</option>
              <option value="finished">Completed</option>
              <option value="pending">Pending</option>
              <option value="processing">Processing</option>
              <option value="cancelled">Cancelled</option>
            </select>
          </div>

          <button
            onClick={resetFilters}
            className="clear-filters-btn"
          >
            Clear filters
          </button>
        </div>

        <div className="history-table">
          {isLoading ? (
            <div className="loading">Loading...</div>
          ) : filteredData.length > 0 ? (
            <>
              <div className="table-header">
                {activeTab === 'transactions' ? (
                  <>
                    <div onClick={() => handleSort('date')}>
                      Date {sortConfig.key === 'date' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div onClick={() => handleSort('type')}>
                      Type {sortConfig.key === 'type' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div onClick={() => handleSort('currency')}>
                      Currency {sortConfig.key === 'currency' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div onClick={() => handleSort('amount')}>
                      Amount {sortConfig.key === 'amount' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div>Fee</div>
                    <div onClick={() => handleSort('status')}>
                      Status {sortConfig.key === 'status' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div>Actions</div>
                  </>
                ) : (
                  <>
                    <div onClick={() => handleSort('date')}>
                      Date {sortConfig.key === 'date' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div>Pair</div>
                    <div onClick={() => handleSort('amount')}>
                      Amount {sortConfig.key === 'amount' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div>Fee</div>
                    <div onClick={() => handleSort('status')}>
                      Status {sortConfig.key === 'status' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div>Completed</div>
                    <div>Actions</div>
                  </>
                )}
              </div>

              <div className="table-body">
                {filteredData.map((item) => (
                  <React.Fragment key={item.id}>
                    <div className="table-row">
                      {activeTab === 'transactions' ? (
                        <>
                          <div>{formatDate(item.date)}</div>
                          <div className={`type-${item.type}`}>{item.type}</div>
                          <div>{item.currency}</div>
                          <div>{item.amount}</div>
                          <div>{item.fee}</div>
                          <div className={`status-${item.status}`}>{item.status}</div>
                          <div className="actions">
                            <button
                              onClick={() => setExpandedId(expandedId === item.id ? null : item.id)}
                              className="details-btn"
                            >
                              {expandedId === item.id ? 'Hide' : 'Details'}
                            </button>
                          </div>
                        </>
                      ) : (
                        <>
                          <div>{formatDate(item.date)}</div>
                          <div>{item.sellCurrency}/{item.buyCurrency}</div>
                          <div>{item.amount}</div>
                          <div>{item.fee}</div>
                          <div className={`status-${item.status}`}>{item.status}</div>
                          <div>{item.completedAmount}/{item.amount}</div>
                          <div className="actions">
                            <button
                              onClick={() => setExpandedId(expandedId === item.id ? null : item.id)}
                              className="details-btn"
                            >
                              {expandedId === item.id ? 'Hide' : 'Details'}
                            </button>
                            {(item.rawStatus === 'pending' || item.rawStatus === 'processing') && (
                              <button
                                onClick={() => handleCancelClick(item.id)}
                                className="cancel-btn"
                              >
                                Cancel
                              </button>
                            )}
                          </div>
                        </>
                      )}
                    </div>

                    {expandedId === item.id && (
                      <div className="expanded-row">
                        {activeTab === 'transactions' ? (
                          <div className="wallet-info">
                            <strong>Wallet:</strong> {item.wallet}
                          </div>
                        ) : (
                          <div className="trade-details">
                            <div><strong>Sell:</strong> {item.sellCurrency}</div>
                            <div><strong>Buy:</strong> {item.buyCurrency}</div>
                            <div><strong>Progress:</strong> {item.amount > 0 ? (item.completedAmount / item.amount * 100).toFixed(2) : 0}%</div>
                          </div>
                        )}
                      </div>
                    )}
                  </React.Fragment>
                ))}
              </div>
            </>
          ) : (
            <div className="no-data">No records found</div>
          )}
        </div>
      </div>

      {/* Модальное окно подтверждения отмены */}
      {cancelModal.show && (
        <div className="modal-overlay">
          <div className="modal">
            <div className="modal-header">
              <h3>Confirm cancellation</h3>
              <button className="modal-close" onClick={closeModal}>×</button>
            </div>
            <div className="modal-content">
              <p>Are you sure you want to cancel this {cancelModal.type === 'trades' ? 'trade' : 'transaction'}?</p>
              <div className="modal-buttons">
                <button
                  onClick={confirmCancel}
                  className="modal-confirm-btn"
                >
                  Yes, cancel
                </button>
                <button
                  onClick={closeModal}
                  className="modal-cancel-btn"
                >
                  No, keep it
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

      {/* Модальное окно ошибки */}
      {errorModal.show && (
        <div className="modal-overlay" onClick={closeErrorModal}>
          <div className="modal" onClick={e => e.stopPropagation()}>
            <div className="modal-header">
              <h3>Error</h3>
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

export default Profile;