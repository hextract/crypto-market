package implementation

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func (ds *DatabaseService) CreateHotWallet(ctx context.Context, currency string) error {
	logEntry := ds.logger.WithFields(logrus.Fields{
		"method":   "CreateHotWallet",
		"currency": currency,
	})

	// Check BTC operations flag
	currency = strings.ToUpper(currency)
	if currency == "BTC" && os.Getenv("ENABLE_BTC_OPERATIONS") != "true" {
		logEntry.Error("BTC operations are disabled")
		return errors.New("BTC operations are currently disabled")
	}

	// Validate currency
	if currency != "USDT" && currency != "BTC" {
		logEntry.Error("Unsupported currency")
		return errors.New("unsupported currency")
	}

	// Start database transaction
	tx, err := ds.pool.Begin(ctx)
	if err != nil {
		logEntry.WithError(err).Error("Failed to start DB transaction")
		return errors.New("internal server error")
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
			_, logErr := ds.pool.Exec(ctx,
				`INSERT INTO transaction_logs (transaction_id, event_type, message, created_at)
				 VALUES ($1, $2, $3, $4)`,
				"", "error", fmt.Sprintf("Create hot wallet failed: %v", err), time.Now())
			if logErr != nil {
				logEntry.WithError(logErr).Error("Failed to log error")
			}
		}
	}()

	// Get currency ID
	var currencyID int
	err = tx.QueryRow(ctx, "SELECT currency_id FROM currencies WHERE name = $1", currency).Scan(&currencyID)
	if err != nil {
		logEntry.WithError(err).Error("Currency not found")
		return errors.New("invalid currency")
	}

	// Check if hot wallet already exists
	var exists bool
	err = tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM hot_wallets WHERE currency_id = $1)", currencyID).Scan(&exists)
	if err != nil {
		logEntry.WithError(err).Error("Failed to check existing hot wallet")
		return errors.New("internal server error")
	}
	if exists {
		logEntry.Info("Hot wallet already exists for currency")
		tx.Rollback(ctx)
		return errors.New("hot wallet already exists")
	}

	// Generate hot wallet address and key
	var address, encryptedKey string
	if currency == "BTC" {
		if os.Getenv("ENABLE_BTC_OPERATIONS") == "true" {
			// Placeholder for BTC node integration
			logEntry.Info("BTC hot wallet generation placeholder")
			return errors.New("BTC node not yet implemented")
		}
		// Placeholder
		address = "bc1_hot_placeholder_" + uuid.New().String()
		encryptedKey = "encrypted_btc_hot_key_placeholder"
	} else { // USDT
		privKey, err := crypto.GenerateKey()
		if err != nil {
			logEntry.WithError(err).Error("Failed to generate private key")
			return errors.New("internal server error")
		}
		address = crypto.PubkeyToAddress(privKey.PublicKey).Hex()
		privKeyHex := hex.EncodeToString(crypto.FromECDSA(privKey))
		encryptedKey, err = ds.cryptoService.Encrypt([]byte(privKeyHex))
		if err != nil {
			logEntry.WithError(err).Error("Failed to encrypt private key")
			return errors.New("internal server error")
		}
	}

	// Save hot wallet
	_, err = tx.Exec(ctx,
		`INSERT INTO hot_wallets (address, encrypted_key, currency_id, created_at)
		 VALUES ($1, $2, $3, $4)`,
		address, encryptedKey, currencyID, time.Now())
	if err != nil {
		logEntry.WithError(err).Error("Failed to save hot wallet")
		return errors.New("internal server error")
	}

	// Commit transaction
	if err = tx.Commit(ctx); err != nil {
		logEntry.WithError(err).Error("Failed to commit DB transaction")
		return errors.New("internal server error")
	}

	logEntry.WithField("address", address).Info("Hot wallet created")
	return nil
}
