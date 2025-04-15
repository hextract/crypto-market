package implementation

import (
	"context"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"math/big"
	"os"
	"strings"
	"time"
)

func (ds *DatabaseService) Withdraw(ctx context.Context, userID, currency, address string, amount float32) (*models.WithdrawResponse, error) {
	logEntry := ds.logger.WithFields(logrus.Fields{
		"method":   "Withdraw",
		"user_id":  userID,
		"currency": currency,
		"amount":   amount,
		"address":  address,
	})

	// Check BTC operations flag
	currency = strings.ToUpper(currency)
	if currency == "BTC" && os.Getenv("ENABLE_BTC_OPERATIONS") != "true" {
		logEntry.Error("BTC operations are disabled")
		return nil, errors.New("BTC operations are currently disabled")
	}

	// Validate inputs
	if userID == "" {
		logEntry.Error("User ID is empty")
		return nil, errors.New("user_id is required")
	}
	if currency != "USDT" && currency != "BTC" {
		logEntry.Error("Unsupported currency")
		return nil, errors.New("unsupported currency")
	}
	if amount <= 0 {
		logEntry.Error("Invalid amount")
		return nil, errors.New("amount must be positive")
	}
	if address == "" {
		logEntry.Error("Address is empty")
		return nil, errors.New("address is required")
	}
	if currency == "USDT" && !common.IsHexAddress(address) {
		logEntry.Error("Invalid Ethereum address")
		return nil, errors.New("invalid address")
	}
	if currency == "BTC" {
		_, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
		if err != nil {
			logEntry.WithError(err).Error("Invalid Bitcoin address")
			return nil, errors.New("invalid address")
		}
	}

	// Start database transaction
	tx, err := ds.pool.Begin(ctx)
	if err != nil {
		logEntry.WithError(err).Error("Failed to start DB transaction")
		return nil, errors.New("internal server error")
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
			_, logErr := ds.pool.Exec(ctx,
				`INSERT INTO transaction_logs (transaction_id, event_type, message, created_at)
				 VALUES ($1, $2, $3, $4)`,
				"", "error", fmt.Sprintf("Withdraw failed: %v", err), time.Now())
			if logErr != nil {
				logEntry.WithError(logErr).Error("Failed to log transaction error")
			}
		}
	}()

	// Get currency ID
	var currencyID int
	err = tx.QueryRow(ctx, "SELECT currency_id FROM currencies WHERE name = $1", currency).Scan(&currencyID)
	if err != nil {
		logEntry.WithError(err).Error("Currency not found")
		return nil, errors.New("invalid currency")
	}

	// Check balance
	var balance float32
	err = tx.QueryRow(ctx,
		`SELECT balance FROM user_balance WHERE user_id = $1 AND currency_id = $2 FOR UPDATE`,
		userID, currencyID).Scan(&balance)
	if err == pgx.ErrNoRows {
		logEntry.Error("Balance not found")
		return nil, errors.New("insufficient balance")
	}
	if err != nil {
		logEntry.WithError(err).Error("Failed to check balance")
		return nil, errors.New("internal server error")
	}
	if balance < amount {
		logEntry.WithFields(logrus.Fields{"balance": balance, "requested": amount}).Error("Insufficient balance")
		return nil, errors.New("insufficient balance")
	}

	// Deduct balance
	_, err = tx.Exec(ctx,
		`UPDATE user_balance SET balance = balance - $1 WHERE user_id = $2 AND currency_id = $3`,
		amount, userID, currencyID)
	if err != nil {
		logEntry.WithError(err).Error("Failed to update balance")
		return nil, errors.New("internal server error")
	}

	// Create transaction
	txID := "tx_with_" + uuid.New().String()
	_, err = tx.Exec(ctx,
		`INSERT INTO transactions (id, user_id, currency_id, type, amount, commission, status, address, created_at)
		 VALUES ($1, $2, $3, 'withdraw', $4, 0, 'pending', $5, $6)`,
		txID, userID, currencyID, amount, address, time.Now())
	if err != nil {
		logEntry.WithError(err).Error("Failed to create transaction")
		return nil, errors.New("internal server error")
	}

	// Commit transaction
	if err = tx.Commit(ctx); err != nil {
		logEntry.WithError(err).Error("Failed to commit DB transaction")
		return nil, errors.New("internal server error")
	}

	// Process withdrawal asynchronously
	go ds.processWithdraw(context.Background(), logEntry, txID, userID, currency, address, amount, currencyID)

	logEntry.WithField("tx_id", txID).Info("Withdraw request created")

	status := "pending"
	response := models.WithdrawResponse{
		ID:     &txID,
		Status: &status,
	}
	return &response, nil
}

// processWithdraw handles the blockchain transaction for withdrawal.
func (ds *DatabaseService) processWithdraw(ctx context.Context, logEntry *logrus.Entry, txID, userID, currency, toAddress string, amount float32, currencyID int) {
	logEntry = logEntry.WithField("tx_id", txID)

	updateStatus := func(status, txHash, errorMsg string) {
		_, err := ds.pool.Exec(ctx,
			`UPDATE transactions SET status = $1, tx_hash = $2 WHERE id = $3`,
			status, txHash, txID)
		if err != nil {
			logEntry.WithError(err).Error("Failed to update transaction status")
		}
		logMsg := fmt.Sprintf("Transaction %s: %s", status, errorMsg)
		if errorMsg == "" {
			logMsg = fmt.Sprintf("Transaction %s", status)
		}
		_, err = ds.pool.Exec(ctx,
			`INSERT INTO transaction_logs (transaction_id, event_type, message, created_at)
			 VALUES ($1, $2, $3, $4)`,
			txID, status, logMsg, time.Now())
		if err != nil {
			logEntry.WithError(err).Error("Failed to log transaction event")
		}
	}

	if currency == "BTC" {
		if os.Getenv("ENABLE_BTC_OPERATIONS") == "true" {
			// Placeholder for BTC node integration
			logEntry.Info("BTC withdrawal placeholder")
			updateStatus("pending", "", "Awaiting BTC node implementation")
			return
		}
		logEntry.Info("BTC withdrawal skipped due to disabled operations")
		updateStatus("cancelled", "", "BTC operations are disabled")
		return
	}

	// USDT withdrawal
	var walletAddress, encryptedKey string
	err := ds.pool.QueryRow(ctx,
		`SELECT address, encrypted_key FROM hot_wallets WHERE currency_id = $1`,
		currencyID).Scan(&walletAddress, &encryptedKey)
	if err != nil {
		logEntry.WithError(err).Error("Failed to fetch hot wallet")
		updateStatus("cancelled", "", "hot wallet not configured")
		// Refund balance
		_, err := ds.pool.Exec(ctx,
			`INSERT INTO user_balance (user_id, currency_id, balance)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (user_id, currency_id)
			 DO UPDATE SET balance = user_balance.balance + $3`,
			userID, currencyID, amount)
		if err != nil {
			logEntry.WithError(err).Error("Failed to refund balance")
		}
		return
	}

	privKeyBytes, err := ds.cryptoService.Decrypt(encryptedKey)
	if err != nil {
		logEntry.WithError(err).Error("Failed to decrypt private key")
		updateStatus("cancelled", "", "internal server error")
		// Refund balance
		_, err := ds.pool.Exec(ctx,
			`INSERT INTO user_balance (user_id, currency_id, balance)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (user_id, currency_id)
			 DO UPDATE SET balance = user_balance.balance + $3`,
			userID, currencyID, amount)
		if err != nil {
			logEntry.WithError(err).Error("Failed to refund balance")
		}
		return
	}
	privKey, err := crypto.HexToECDSA(string(privKeyBytes))
	if err != nil {
		logEntry.WithError(err).Error("Invalid private key")
		updateStatus("cancelled", "", "internal server error")
		// Refund balance
		_, err := ds.pool.Exec(ctx,
			`INSERT INTO user_balance (user_id, currency_id, balance)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (user_id, currency_id)
			 DO UPDATE SET balance = user_balance.balance + $3`,
			userID, currencyID, amount)
		if err != nil {
			logEntry.WithError(err).Error("Failed to refund balance")
		}
		return
	}

	nonce, err := ds.blockchainClient.EthClient.PendingNonceAt(ctx, common.HexToAddress(walletAddress))
	if err != nil {
		logEntry.WithError(err).Error("Failed to get nonce")
		updateStatus("cancelled", "", "failed to prepare transaction")
		// Refund balance
		_, err := ds.pool.Exec(ctx,
			`INSERT INTO user_balance (user_id, currency_id, balance)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (user_id, currency_id)
			 DO UPDATE SET balance = user_balance.balance + $3`,
			userID, currencyID, amount)
		if err != nil {
			logEntry.WithError(err).Error("Failed to refund balance")
		}
		return
	}
	gasPrice, err := ds.blockchainClient.EthClient.SuggestGasPrice(ctx)
	if err != nil {
		logEntry.WithError(err).Error("Failed to get gas price")
		updateStatus("cancelled", "", "failed to prepare transaction")
		// Refund balance
		_, err := ds.pool.Exec(ctx,
			`INSERT INTO user_balance (user_id, currency_id, balance)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (user_id, currency_id)
			 DO UPDATE SET balance = user_balance.balance + $3`,
			userID, currencyID, amount)
		if err != nil {
			logEntry.WithError(err).Error("Failed to refund balance")
		}
		return
	}

	amountBig := big.NewFloat(float64(amount * 1e6)) // USDT has 6 decimals
	amountInt, _ := amountBig.Int(nil)
	usdtAddress := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7") // USDT ERC-20 contract
	usdtABI, err := abi.JSON(strings.NewReader(`[
		{
			"constant": false,
			"inputs": [
				{"name": "_to", "type": "address"},
				{"name": "_value", "type": "uint256"}
			],
			"name": "transfer",
			"outputs": [{"name": "", "type": "bool"}],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		}
	]`))
	if err != nil {
		logEntry.WithError(err).Error("Failed to parse USDT ABI")
		updateStatus("cancelled", "", "internal server error")
		// Refund balance
		_, err := ds.pool.Exec(ctx,
			`INSERT INTO user_balance (user_id, currency_id, balance)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (user_id, currency_id)
			 DO UPDATE SET balance = user_balance.balance + $3`,
			userID, currencyID, amount)
		if err != nil {
			logEntry.WithError(err).Error("Failed to refund balance")
		}
		return
	}
	usdtContract := bind.NewBoundContract(
		usdtAddress,
		usdtABI,
		ds.blockchainClient.EthClient,
		ds.blockchainClient.EthClient,
		ds.blockchainClient.EthClient,
	)

	txData, err := usdtContract.Transact(&bind.TransactOpts{
		From:     common.HexToAddress(walletAddress),
		Nonce:    big.NewInt(int64(nonce)),
		GasPrice: gasPrice,
		GasLimit: 60000,
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, types.HomesteadSigner{}, privKey)
		},
		Context: ctx,
	}, "transfer", common.HexToAddress(toAddress), amountInt)
	if err != nil {
		logEntry.WithError(err).Error("Failed to create transaction")
		updateStatus("cancelled", "", "failed to create transaction")
		// Refund balance
		_, err := ds.pool.Exec(ctx,
			`INSERT INTO user_balance (user_id, currency_id, balance)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (user_id, currency_id)
			 DO UPDATE SET balance = user_balance.balance + $3`,
			userID, currencyID, amount)
		if err != nil {
			logEntry.WithError(err).Error("Failed to refund balance")
		}
		return
	}

	if err = ds.blockchainClient.EthClient.SendTransaction(ctx, txData); err != nil {
		logEntry.WithError(err).Error("Failed to send transaction")
		updateStatus("cancelled", "", "failed to send transaction")
		// Refund balance
		_, err := ds.pool.Exec(ctx,
			`INSERT INTO user_balance (user_id, currency_id, balance)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (user_id, currency_id)
			 DO UPDATE SET balance = user_balance.balance + $3`,
			userID, currencyID, amount)
		if err != nil {
			logEntry.WithError(err).Error("Failed to refund balance")
		}
		return
	}

	go func() {
		for i := 0; i < 60; i++ {
			receipt, err := ds.blockchainClient.EthClient.TransactionReceipt(ctx, txData.Hash())
			if err == nil && receipt != nil {
				if receipt.Status == types.ReceiptStatusSuccessful {
					updateStatus("finished", txData.Hash().Hex(), "")
					logEntry.WithField("tx_hash", txData.Hash().Hex()).Info("Withdrawal completed")
				} else {
					updateStatus("cancelled", txData.Hash().Hex(), "transaction failed")
					logEntry.WithField("tx_hash", txData.Hash().Hex()).Error("Transaction failed")
					// Refund balance
					_, err := ds.pool.Exec(ctx,
						`INSERT INTO user_balance (user_id, currency_id, balance)
						 VALUES ($1, $2, $3)
						 ON CONFLICT (user_id, currency_id)
						 DO UPDATE SET balance = user_balance.balance + $3`,
						userID, currencyID, amount)
					if err != nil {
						logEntry.WithError(err).Error("Failed to refund balance")
					}
				}
				return
			}
			time.Sleep(10 * time.Second)
		}
		logEntry.Error("Transaction timed out")
		updateStatus("cancelled", txData.Hash().Hex(), "transaction timed out")
		// Refund balance
		_, err := ds.pool.Exec(ctx,
			`INSERT INTO user_balance (user_id, currency_id, balance)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (user_id, currency_id)
			 DO UPDATE SET balance = user_balance.balance + $3`,
			userID, currencyID, amount)
		if err != nil {
			logEntry.WithError(err).Error("Failed to refund balance")
		}
	}()
}
