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
import { useTranslation } from 'react-i18next';

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
  const { t } = useTranslation();

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
      showError(t('profile.modals.errorLoading'));
    } finally {
      setIsLoading(false);
    }
  }, [activeTab, filters, t]);

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
    type: mapOperation(item.operation),
    type_style: item.operation,
    currency: item.currency,
    amount: item.amount,
    fee: item.commission || 0,
    status: mapTradeStatus(item.status),
    status_style: item.status,
    wallet: item.address,
    rawStatus: item.status
  });

  const mapTrade = (item) => ({
    id: item.id,
    date: new Date(item.create_date),
    cancelledDate: item.complete_date ? new Date(item.complete_date) : null,
    sellCurrency: item.from_currency,
    buyCurrency: item.to_currency,
    amount: item.amount_to_buy,
    fee: item.commission || 0,
    status: mapTradeStatus(item.status),
    status_style: item.status,
    completedAmount: item.bought_amount || 0,
    rawStatus: item.status
  });

  const mapOperation = (op) => {
    switch (op) {
      case 'withdraw': return t('profile.withdraw');
      case 'deposit': return t('profile.deposit');
      default: return op;
    }
  };

  const mapTradeStatus = (status) => {
    switch (status) {
      case 'finished': return t('profile.filters.statuses.completed');
      case 'processing': return t('profile.filters.statuses.processing');
      case 'pending': return t('profile.filters.statuses.pending');
      case 'cancelled': return t('profile.filters.statuses.cancelled');
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

    return !(filters.status !== 'all' && item.rawStatus !== filters.status);
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
      await loadData();
      setCancelModal({ show: false, id: null, type: null });
    } catch (error) {
      console.error('Failed to cancel trade:', error);
      showError(t('profile.modals.errorCancel'));
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
          <span>{t('navbar.logo')}</span>
        </div>
        <div className="nav-links">
          <a href="/main">{t('navbar.trade')}</a>
          <a href="/profile" className="active">{t('navbar.profile')}</a>
          <button onClick={handleLogout} className="logout-btn">{t('navbar.logout')}</button>
        </div>
      </nav>

      <div className="profile-content">
        <div className="profile-tabs">
          <button
            className={activeTab === 'trades' ? 'active' : ''}
            onClick={() => setActiveTab('trades')}
          >
            {t('profile.tabs.trades')}
          </button>
          <button
            className={activeTab === 'transactions' ? 'active' : ''}
            onClick={() => setActiveTab('transactions')}
          >
            {t('profile.tabs.transactions')}
          </button>
        </div>

        <div className="filters">
          <div className="filter-group">
            <label>{t('profile.filters.time')}</label>
            <select
              value={filters.time}
              onChange={(e) => handleFilterChange('time', e.target.value)}
            >
              <option value="all">{t('profile.filters.allTime')}</option>
              <option value="hour">{t('profile.filters.lastHour')}</option>
              <option value="day">{t('profile.filters.lastDay')}</option>
              <option value="week">{t('profile.filters.lastWeek')}</option>
              <option value="month">{t('profile.filters.lastMonth')}</option>
              <option value="year">{t('profile.filters.lastYear')}</option>
            </select>
          </div>

          {activeTab === 'transactions' && (
            <div className="filter-group">
              <label>{t('profile.filters.type')}</label>
              <select
                value={filters.type}
                onChange={(e) => handleFilterChange('type', e.target.value)}
              >
                <option value="all">{t('profile.filters.allTypes')}</option>
                <option value="deposit">{t('profile.filters.deposits')}</option>
                <option value="withdraw">{t('profile.filters.withdrawals')}</option>
              </select>
            </div>
          )}

          <div className="filter-group">
            <label>{t('profile.filters.status')}</label>
            <select
              value={filters.status}
              onChange={(e) => handleFilterChange('status', e.target.value)}
            >
              <option value="all">{t('profile.filters.allStatuses')}</option>
              <option value="finished">{t('profile.filters.statuses.completed')}</option>
              <option value="pending">{t('profile.filters.statuses.pending')}</option>
              <option value="processing">{t('profile.filters.statuses.processing')}</option>
              <option value="cancelled">{t('profile.filters.statuses.cancelled')}</option>
            </select>
          </div>

          <button
            onClick={resetFilters}
            className="clear-filters-btn"
          >
            {t('profile.filters.clear')}
          </button>
        </div>

        <div className="history-table">
          {isLoading ? (
            <div className="loading">{t('profile.table.loading')}</div>
          ) : filteredData.length > 0 ? (
            <>
              <div className="table-header">
                {activeTab === 'transactions' ? (
                  <>
                    <div onClick={() => handleSort('date')}>
                      {t('profile.table.date')} {sortConfig.key === 'date' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div onClick={() => handleSort('type')}>
                      {t('profile.table.type')} {sortConfig.key === 'type' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div onClick={() => handleSort('currency')}>
                      {t('profile.table.currency')} {sortConfig.key === 'currency' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div onClick={() => handleSort('amount')}>
                      {t('profile.table.amount')} {sortConfig.key === 'amount' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div>{t('profile.table.fee')}</div>
                    <div onClick={() => handleSort('status')}>
                      {t('profile.table.status')} {sortConfig.key === 'status' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div>{t('profile.table.actions')}</div>
                  </>
                ) : (
                  <>
                    <div onClick={() => handleSort('date')}>
                      {t('profile.table.date')} {sortConfig.key === 'date' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div>{t('profile.table.pair')}</div>
                    <div onClick={() => handleSort('amount')}>
                      {t('profile.table.amount')} {sortConfig.key === 'amount' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div>{t('profile.table.fee')}</div>
                    <div onClick={() => handleSort('status')}>
                      {t('profile.table.status')} {sortConfig.key === 'status' && (sortConfig.direction === 'asc' ? '↑' : '↓')}
                    </div>
                    <div>{t('profile.table.completed')}</div>
                    <div>{t('profile.table.actions')}</div>
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
                          <div className={`type-${item.type_style}`}>{item.type}</div>
                          <div>{item.currency}</div>
                          <div>{item.amount}</div>
                          <div>{item.fee}</div>
                          <div className={`status-${item.status_style}`}>{item.status}</div>
                          <div className="actions">
                            <button
                              onClick={() => setExpandedId(expandedId === item.id ? null : item.id)}
                              className="details-btn"
                            >
                              {expandedId === item.id ? t('profile.table.hide') : t('profile.table.details')}
                            </button>
                          </div>
                        </>
                      ) : (
                        <>
                          <div>{formatDate(item.date)}</div>
                          <div>{item.sellCurrency}/{item.buyCurrency}</div>
                          <div>{item.amount}</div>
                          <div>{item.fee}</div>
                          <div className={`status-${item.status_style}`}>{item.status}</div>
                          <div>{item.completedAmount}/{item.amount}</div>
                          <div className="actions">
                            <button
                              onClick={() => setExpandedId(expandedId === item.id ? null : item.id)}
                              className="details-btn"
                            >
                              {expandedId === item.id ? t('profile.table.hide') : t('profile.table.details')}
                            </button>
                            {(item.rawStatus === 'pending' || item.rawStatus === 'processing') && (
                              <button
                                onClick={() => handleCancelClick(item.id)}
                                className="cancel-btn"
                              >
                                {t('profile.table.cancel')}
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
                            <strong>{t('profile.table.wallet')}</strong> {item.wallet}
                          </div>
                        ) : (
                          <div className="trade-details">
                            <div><strong>{t('profile.table.sell')}</strong> {item.sellCurrency}</div>
                            <div><strong>{t('profile.table.buy')}</strong> {item.buyCurrency}</div>
                            <div><strong>{t('profile.table.progress')}</strong> {item.amount > 0 ? (item.completedAmount / item.amount * 100).toFixed(2) : 0}%</div>
                            {item.cancelledDate && item.cancelledDate.getFullYear() > 1 && (
                              <div><strong>{t('profile.table.completedDate')}</strong> {formatDate(item.cancelledDate)}</div>
                            )}
                          </div>
                        )}
                      </div>
                    )}
                  </React.Fragment>
                ))}
              </div>
            </>
          ) : (
            <div className="no-data">{t('profile.table.noData')}</div>
          )}
        </div>
      </div>

      {cancelModal.show && (
        <div className="modal-overlay">
          <div className="modal">
            <div className="modal-header">
              <h3>{t('profile.modals.cancelTitle')}</h3>
              <button className="modal-close" onClick={closeModal}>×</button>
            </div>
            <div className="modal-content">
              <p>{t('profile.modals.cancelMessage', { type: cancelModal.type === 'trades' ? t('profile.tabs.trades') : t('profile.tabs.transactions') })}</p>
              <div className="modal-buttons">
                <button
                  onClick={confirmCancel}
                  className="modal-confirm-btn"
                >
                  {t('profile.modals.cancelConfirm')}
                </button>
                <button
                  onClick={closeModal}
                  className="modal-cancel-btn"
                >
                  {t('profile.modals.cancelKeep')}
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

      {errorModal.show && (
        <div className="modal-overlay" onClick={closeErrorModal}>
          <div className="modal" onClick={e => e.stopPropagation()}>
            <div className="modal-header">
              <h3>{t('profile.modals.errorTitle')}</h3>
              <button className="modal-close" onClick={closeErrorModal}>×</button>
            </div>
            <div className="modal-content">
              <p>{errorModal.message}</p>
              <div className="modal-buttons">
                <button
                  onClick={closeErrorModal}
                  className="modal-confirm-btn"
                >
                  {t('common.ok')}
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