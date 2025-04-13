CREATE TYPE transaction_type AS ENUM ('deposit', 'withdraw');
CREATE TYPE transaction_status AS ENUM ('pending', 'processing', 'finished', 'cancelled');

CREATE TABLE currencies
(
    currency_id SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    decimals    INT  NOT NULL CHECK (decimals >= 0),
    network     TEXT NOT NULL,
    CONSTRAINT uq_currency_name UNIQUE (name)
);

COMMENT ON TABLE currencies IS 'Stores cryptocurrency metadata';
COMMENT ON COLUMN currencies.decimals IS 'Number of decimal places (e.g., 6 for USDT, 8 for BTC)';
COMMENT ON COLUMN currencies.network IS 'Blockchain network (e.g., Ethereum, Bitcoin)';

CREATE TABLE user_balance
(
    user_id     TEXT           NOT NULL,
    currency_id INT            NOT NULL REFERENCES currencies (currency_id) ON DELETE RESTRICT,
    balance     NUMERIC(20, 8) NOT NULL CHECK (balance >= 0),
    PRIMARY KEY (user_id, currency_id)
);

COMMENT ON TABLE user_balance IS 'Stores user balances per currency';
COMMENT ON COLUMN user_balance.balance IS 'User balance in the specified currency';

CREATE TABLE transactions
(
    id            TEXT PRIMARY KEY,
    user_id       TEXT                     NOT NULL,
    currency_id   INT                      NOT NULL REFERENCES currencies (currency_id) ON DELETE RESTRICT,
    type          transaction_type         NOT NULL,
    amount        NUMERIC(20, 8)           NOT NULL CHECK (amount > 0),
    commission    NUMERIC(20, 8)           NOT NULL DEFAULT 0 CHECK (commission >= 0),
    status        transaction_status       NOT NULL DEFAULT 'pending',
    address       TEXT,
    tx_hash       TEXT,
    encrypted_key TEXT,
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE transactions IS 'Stores deposit and withdrawal transactions';
COMMENT ON COLUMN transactions.id IS 'Unique transaction ID (e.g., tx_dep_123)';
COMMENT ON COLUMN transactions.address IS 'Recipient address for withdrawals, deposit address for deposits';
COMMENT ON COLUMN transactions.tx_hash IS 'Blockchain transaction hash';
COMMENT ON COLUMN transactions.encrypted_key IS 'Encrypted private key for deposit address';

CREATE TABLE currency_rate
(
    from_id    INT                      NOT NULL REFERENCES currencies (currency_id) ON DELETE RESTRICT,
    to_id      INT                      NOT NULL REFERENCES currencies (currency_id) ON DELETE RESTRICT,
    cost       NUMERIC(20, 8)           NOT NULL CHECK (cost > 0),
    valid_from TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (from_id, to_id, valid_from)
);

COMMENT ON TABLE currency_rate IS 'Stores exchange rates between currencies';
COMMENT ON COLUMN currency_rate.cost IS 'Exchange rate from one currency to another';

CREATE TABLE transaction_logs
(
    id             SERIAL PRIMARY KEY,
    transaction_id TEXT                     REFERENCES transactions (id) ON DELETE SET NULL,
    event_type     TEXT                     NOT NULL,
    message        TEXT                     NOT NULL,
    created_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE transaction_logs IS 'Stores logs for transaction events (e.g., errors)';
COMMENT ON COLUMN transaction_logs.event_type IS 'Type of event (e.g., error, info)';
COMMENT ON COLUMN transaction_logs.message IS 'Event description';

CREATE INDEX idx_transactions_user_id ON transactions (user_id);
CREATE INDEX idx_transactions_currency_id ON transactions (currency_id);
CREATE INDEX idx_transactions_status ON transactions (status);
CREATE INDEX idx_transaction_logs_transaction_id ON transaction_logs (transaction_id);

DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM currencies WHERE name = 'USDT') THEN
            INSERT INTO currencies (name, decimals, network) VALUES ('USDT', 6, 'Ethereum');
        END IF;
        IF NOT EXISTS (SELECT 1 FROM currencies WHERE name = 'BTC') THEN
            INSERT INTO currencies (name, decimals, network) VALUES ('BTC', 8, 'Bitcoin');
        END IF;
    END
$$;


-- NOT EDITED

create table if not exists bids
(
    user_id       text      not null,
    currency_id   int       not null references currencies (currency_id),
    status        text  default 'created',
    create_date   timestamp not null,
    complete_date timestamp,
    min_price     float     not null,
    max_price     float     not null,
    amount_to_buy float     not null,
    bought_amount float default 0.0,
    buy_speed     float,
    avg_price     float
);