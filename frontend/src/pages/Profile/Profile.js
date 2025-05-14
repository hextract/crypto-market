import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { logout } from '../../api/authService';
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

  // Заглушки для данных
  const transactions = [
    {
      id: 1,
      date: new Date('2023-05-15T14:30:00'),
      type: 'deposit',
      currency: 'USDT',
      amount: 1000,
      fee: 10,
      status: 'completed',
      wallet: '0x71C7656EC7ab88b098defB751B7401B5f6d8976F'
    },
    {
      id: 2,
      date: new Date('2023-05-14T10:15:00'),
      type: 'withdraw',
      currency: 'BTC',
      amount: 0.1,
      fee: 0.001,
      status: 'pending',
      wallet: '3FZbgi29cpjq2GjdwV8eyHuJJnkLtktZc5'
    }
  ];

  const trades = [
    {
      id: 1,
      date: new Date('2023-05-15T09:45:00'),
      sellCurrency: 'USDT',
      buyCurrency: 'BTC',
      amount: 500,
      fee: 5,
      status: 'partial',
      completedAmount: 300
    },
    {
      id: 2,
      date: new Date('2023-05-14T16:20:00'),
      sellCurrency: 'BTC',
      buyCurrency: 'USDT',
      amount: 0.05,
      fee: 0.0005,
      status: 'pending',
      completedAmount: 0
    }
  ];

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

    return !(filters.status !== 'all' && item.status !== filters.status);


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

  const confirmCancel = () => {
    console.log(`Canceling ${cancelModal.type} with id: ${cancelModal.id}`);
    // Здесь будет логика отмены заявки
    setCancelModal({ show: false, id: null, type: null });
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
              <option value="completed">Completed</option>
              <option value="pending">Pending</option>
              <option value="partial">Partial</option>
              <option value="failed">Failed</option>
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
            {filteredData.length > 0 ? (
              filteredData.map((item) => (
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
                          {(item.status === 'pending' || item.status === 'partial') && (
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
                          <div><strong>Progress:</strong> {(item.completedAmount / item.amount * 100).toFixed(2)}%</div>
                        </div>
                      )}
                    </div>
                  )}
                </React.Fragment>
              ))
            ) : (
              <div className="no-data">No records found</div>
            )}
          </div>
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
    </div>
  );
};

export default Profile;