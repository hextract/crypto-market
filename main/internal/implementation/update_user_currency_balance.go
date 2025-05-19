package implementation

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (ds *DatabaseService) UpdateUserCurrencyBalance(orderID, currency string, amount float32) error {
	if orderID == "" {
		return errors.New("order ID is required")
	}
	if currency == "" {
		return errors.New("currency is required")
	}

	logEntry := ds.logger.WithFields(logrus.Fields{
		"method":   "UpdateUserCurrencyBalance",
		"order_id": orderID,
		"currency": currency,
		"amount":   amount,
	})

	// Начинаем транзакцию
	tx, err := ds.pool.Begin(context.Background())
	if err != nil {
		logEntry.WithError(err).Error("Failed to begin transaction")
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(context.Background()); rbErr != nil {
				logEntry.WithError(rbErr).Error("Transaction rollback failed")
			}
		}
	}()

	// 1. Получаем user_id из заказа
	var userID string
	err = tx.QueryRow(context.Background(),
		`SELECT user_id FROM bids WHERE id = $1`,
		orderID).Scan(&userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			logEntry.Error("Order not found")
			return errors.New("order not found")
		}
		logEntry.WithError(err).Error("Failed to get order")
		return fmt.Errorf("failed to get order: %w", err)
	}

	// 2. Получаем currency_id
	var currencyID int
	err = tx.QueryRow(context.Background(),
		`SELECT currency_id FROM currencies WHERE name = $1`,
		currency).Scan(&currencyID)
	if err != nil {
		if err == pgx.ErrNoRows {
			logEntry.Error("Currency not found")
			return errors.New("currency not found")
		}
		logEntry.WithError(err).Error("Failed to get currency")
		return fmt.Errorf("failed to get currency: %w", err)
	}

	// 3. Обновляем баланс (увеличиваем или уменьшаем)
	query := `
		INSERT INTO user_balance (user_id, currency_id, balance)
        VALUES($1, $2, 0)
        ON CONFLICT (user_id, currency_id) DO NOTHING;
	`
	rows, insertErr := tx.Query(context.Background(), query, userID, currencyID)
	if insertErr != nil {
		return insertErr
	}
	rows.Close()

	query = `
		UPDATE user_balance
        SET balance = balance + $3
        WHERE user_id = $1 AND currency_id = $2
        RETURNING balance;
	`

	var newBalance float32
	err = tx.QueryRow(context.Background(), query,
		userID, currencyID, amount).Scan(&newBalance)
	if err != nil {
		logEntry.WithError(err).Error("Failed to update balance")
		return fmt.Errorf("failed to update balance: %w", err)
	}

	// Фиксируем транзакцию
	if err = tx.Commit(context.Background()); err != nil {
		logEntry.WithError(err).Error("Failed to commit transaction")
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	logEntry.WithFields(logrus.Fields{
		"user_id":     userID,
		"new_balance": newBalance,
	}).Info("Balance updated successfully")

	return nil
}
