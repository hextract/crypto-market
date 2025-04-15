package implementation

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/sirupsen/logrus"
	"math/big"
	"os"
	"strings"
	"time"
)

func (ds *DatabaseService) Deposit(ctx context.Context, userID, currency string, amount float32) (*models.DepositResponse, error) {
	logEntry := ds.logger.WithFields(logrus.Fields{
		"method":   "Deposit",
		"user_id":  userID,
		"currency": currency,
		"amount":   amount,
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
				"", "error", fmt.Sprintf("Deposit failed: %v", err), time.Now())
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

	// Generate deposit address and key
	var depositAddress, encryptedKey string
	if currency == "BTC" {
		if os.Getenv("ENABLE_BTC_OPERATIONS") == "true" {
			// Placeholder for BTC node integration
			logEntry.Info("BTC deposit address generation placeholder")
			return nil, errors.New("BTC node not yet implemented")
		}
		// Placeholder
		depositAddress = "bc1_placeholder_" + uuid.New().String()
		encryptedKey = "encrypted_btc_key_placeholder"
	} else { // USDT
		privKey, err := crypto.GenerateKey()
		if err != nil {
			logEntry.WithError(err).Error("Failed to generate private key")
			return nil, errors.New("internal server error")
		}
		depositAddress = crypto.PubkeyToAddress(privKey.PublicKey).Hex()
		privKeyHex := hex.EncodeToString(crypto.FromECDSA(privKey))
		encryptedKey, err = ds.cryptoService.Encrypt([]byte(privKeyHex))
		if err != nil {
			logEntry.WithError(err).Error("Failed to encrypt private key")
			return nil, errors.New("internal server error")
		}
	}

	// Save wallet
	walletID := "wallet_" + uuid.New().String()
	_, err = tx.Exec(ctx,
		`INSERT INTO wallets (id, user_id, currency_id, address, encrypted_key, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		walletID, userID, currencyID, depositAddress, encryptedKey, time.Now())
	if err != nil {
		logEntry.WithError(err).Error("Failed to save wallet")
		return nil, errors.New("internal server error")
	}

	// Create transaction
	txID := "tx_dep_" + uuid.New().String()
	_, err = tx.Exec(ctx,
		`INSERT INTO transactions (id, user_id, currency_id, type, amount, commission, status, address, encrypted_key, created_at)
		 VALUES ($1, $2, $3, 'deposit', $4, 0, 'pending', $5, $6, $7)`,
		txID, userID, currencyID, amount, depositAddress, encryptedKey, time.Now())
	if err != nil {
		logEntry.WithError(err).Error("Failed to create transaction")
		return nil, errors.New("internal server error")
	}

	// Commit transaction
	if err = tx.Commit(ctx); err != nil {
		logEntry.WithError(err).Error("Failed to commit DB transaction")
		return nil, errors.New("internal server error")
	}

	// Start async deposit monitoring
	go ds.monitorDeposit(context.Background(), logEntry, txID, userID, currency, depositAddress, amount, currencyID)

	status := "pending"

	logEntry.WithField("tx_id", txID).Info("Deposit request created")

	response := models.DepositResponse{
		ID:      &txID,
		Status:  &status,
		Address: &depositAddress,
	}

	return &response, nil
}

// monitorDeposit monitors the blockchain for incoming deposits and transfers them to the hot wallet.
func (ds *DatabaseService) monitorDeposit(ctx context.Context, logEntry *logrus.Entry, txID, userID, currency, depositAddress string, requestedAmount float32, currencyID int) {
	logEntry = logEntry.WithField("tx_id", txID)

	updateStatus := func(status, txHash, errorMsg string, actualAmount float32) {
		_, err := ds.pool.Exec(ctx,
			`UPDATE transactions SET status = $1, tx_hash = $2, amount = $3 WHERE id = $5`,
			status, txHash, actualAmount, txID)
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
			logEntry.Info("BTC deposit monitoring placeholder")
			updateStatus("pending", "", "Awaiting BTC node implementation", 0)
			return
		}
		logEntry.Info("BTC deposit monitoring skipped due to disabled operations")
		updateStatus("cancelled", "", "BTC operations are disabled", 0)
		return
	}

	// USDT deposit monitoring
	go func() {
		// USDT contract and Transfer event signature
		usdtAddress := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
		transferSig := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")).Hex()
		usdtABI, err := abi.JSON(strings.NewReader(`[
			{
				"anonymous": false,
				"inputs": [
					{"indexed": true, "name": "from", "type": "address"},
					{"indexed": true, "name": "to", "type": "address"},
					{"indexed": false, "name": "value", "type": "uint256"}
				],
				"name": "Transfer",
				"type": "event"
			},
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
			updateStatus("cancelled", "", "internal server error", 0)
			return
		}

		// Get hot wallet address
		var hotWalletAddress string
		err = ds.pool.QueryRow(ctx,
			`SELECT address FROM hot_wallets WHERE currency_id = $1`,
			currencyID).Scan(&hotWalletAddress)
		if err != nil {
			logEntry.WithError(err).Error("Failed to fetch hot wallet")
			updateStatus("cancelled", "", "hot wallet not configured", 0)
			return
		}

		// Get deposit wallet private key
		var encryptedKey string
		err = ds.pool.QueryRow(ctx,
			`SELECT encrypted_key FROM wallets WHERE address = $1 AND currency_id = $2`,
			depositAddress, currencyID).Scan(&encryptedKey)
		if err != nil {
			logEntry.WithError(err).Error("Failed to fetch deposit wallet key")
			updateStatus("cancelled", "", "internal server error", 0)
			return
		}
		privKeyBytes, err := ds.cryptoService.Decrypt(encryptedKey)
		if err != nil {
			logEntry.WithError(err).Error("Failed to decrypt private key")
			updateStatus("cancelled", "", "internal server error", 0)
			return
		}
		privKey, err := crypto.HexToECDSA(string(privKeyBytes))
		if err != nil {
			logEntry.WithError(err).Error("Invalid private key")
			updateStatus("cancelled", "", "internal server error", 0)
			return
		}

		// Monitor for Transfer events
		for i := 0; i < 360; i++ { // Monitor for 1 hour
			latestBlock, err := ds.blockchainClient.EthClient.BlockNumber(ctx)
			if err != nil {
				logEntry.WithError(err).Error("Failed to get latest block")
				time.Sleep(10 * time.Second)
				continue
			}
			fromBlock := latestBlock - 100 // Check last ~15 minutes
			if fromBlock < 0 {
				fromBlock = 0
			}

			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(fromBlock)),
				ToBlock:   big.NewInt(int64(latestBlock)),
				Addresses: []common.Address{usdtAddress},
				Topics: [][]common.Hash{
					{common.HexToHash(transferSig)},
					{}, // from: any address
					{common.BytesToHash(common.HexToAddress(depositAddress).Bytes())}, // to: depositAddress
				},
			}
			logs, err := ds.blockchainClient.EthClient.FilterLogs(ctx, query)
			if err != nil {
				logEntry.WithError(err).Error("Failed to fetch logs")
				time.Sleep(10 * time.Second)
				continue
			}

			for _, log := range logs {
				if len(log.Topics) < 3 {
					continue
				}
				var transferEvent struct {
					From  common.Address
					To    common.Address
					Value *big.Int
				}
				err = usdtABI.UnpackIntoInterface(&transferEvent, "Transfer", log.Data)
				if err != nil {
					logEntry.WithError(err).Error("Failed to unpack transfer event")
					continue
				}
				transferEvent.From = common.HexToAddress(log.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(log.Topics[2].Hex())

				if transferEvent.To.Hex() != depositAddress {
					continue
				}

				// Convert amount (USDT has 6 decimals)
				amountBig := new(big.Float).SetInt(transferEvent.Value)
				actualAmount, _ := new(big.Float).Quo(amountBig, big.NewFloat(1e6)).Float32()
				if actualAmount < requestedAmount {
					logEntry.WithFields(logrus.Fields{
						"actual_amount":    actualAmount,
						"requested_amount": requestedAmount,
					}).Error("Deposit amount too low")
					updateStatus("cancelled", "", "deposit amount below requested", actualAmount)
					return
				}

				// Transfer to hot wallet
				usdtContract := bind.NewBoundContract(
					usdtAddress,
					usdtABI,
					ds.blockchainClient.EthClient,
					ds.blockchainClient.EthClient,
					ds.blockchainClient.EthClient,
				)
				nonce, err := ds.blockchainClient.EthClient.PendingNonceAt(ctx, common.HexToAddress(depositAddress))
				if err != nil {
					logEntry.WithError(err).Error("Failed to get nonce")
					updateStatus("cancelled", "", "failed to prepare transaction", actualAmount)
					return
				}
				gasPrice, err := ds.blockchainClient.EthClient.SuggestGasPrice(ctx)
				if err != nil {
					logEntry.WithError(err).Error("Failed to get gas price")
					updateStatus("cancelled", "", "failed to prepare transaction", actualAmount)
					return
				}

				txData, err := usdtContract.Transact(&bind.TransactOpts{
					From:     common.HexToAddress(depositAddress),
					Nonce:    big.NewInt(int64(nonce)),
					GasPrice: gasPrice,
					GasLimit: 60000,
					Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
						return types.SignTx(tx, types.HomesteadSigner{}, privKey)
					},
					Context: ctx,
				}, "transfer", common.HexToAddress(hotWalletAddress), transferEvent.Value)
				if err != nil {
					logEntry.WithError(err).Error("Failed to create hot wallet transfer")
					updateStatus("cancelled", "", "failed to transfer to hot wallet", actualAmount)
					return
				}
				if err = ds.blockchainClient.EthClient.SendTransaction(ctx, txData); err != nil {
					logEntry.WithError(err).Error("Failed to send hot wallet transfer")
					updateStatus("cancelled", "", "failed to transfer to hot wallet", actualAmount)
					return
				}

				// Monitor transfer confirmation
				var confirmed bool
				var transferTxHash = txData.Hash().Hex()
				for j := 0; j < 60; j++ {
					receipt, err := ds.blockchainClient.EthClient.TransactionReceipt(ctx, txData.Hash())
					if err == nil && receipt != nil {
						if receipt.Status == types.ReceiptStatusSuccessful {
							confirmed = true
							break
						}
						logEntry.WithField("tx_hash", transferTxHash).Error("Hot wallet transfer failed")
						updateStatus("cancelled", transferTxHash, "hot wallet transfer failed", actualAmount)
						return
					}
					time.Sleep(10 * time.Second)
				}
				if !confirmed {
					logEntry.WithField("tx_hash", transferTxHash).Error("Hot wallet transfer timed out")
					updateStatus("cancelled", transferTxHash, "hot wallet transfer timed out", actualAmount)
					return
				}

				// Update user balance
				_, err = ds.pool.Exec(ctx,
					`INSERT INTO user_balance (user_id, currency_id, balance)
					 VALUES ($1, $2, $3)
					 ON CONFLICT (user_id, currency_id)
					 DO UPDATE SET balance = user_balance.balance + $3`,
					userID, currencyID, actualAmount)
				if err != nil {
					logEntry.WithError(err).Error("Failed to update balance")
					updateStatus("cancelled", transferTxHash, "failed to update balance", actualAmount)
					return
				}

				// Update transaction
				updateStatus("finished", transferTxHash, "", actualAmount)
				logEntry.WithFields(logrus.Fields{
					"amount":     actualAmount,
					"tx_hash":    transferTxHash,
					"hot_wallet": hotWalletAddress,
				}).Info("Deposit detected and transferred to hot wallet")
				return
			}
			time.Sleep(10 * time.Second)
		}
		logEntry.Info("Deposit monitoring timed out")
		updateStatus("cancelled", "", "deposit not detected", 0)
	}()
}
