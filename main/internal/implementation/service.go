package implementation

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type CryptoService interface {
	Encrypt(data []byte) (string, error)
	Decrypt(encrypted string) ([]byte, error)
}

type DatabaseService struct {
	pool             *pgxpool.Pool
	blockchainClient *BlockchainClient
	cryptoService    CryptoService
	logger           *logrus.Logger
}

func NewDatabaseService(connStr string, blockchainClient *BlockchainClient) (*DatabaseService, error) {
	result := new(DatabaseService)
	newPool, errPool := pgxpool.New(context.Background(), connStr)
	if (errPool != nil) || (newPool == nil) {
		return nil, errPool
	}

	cryptoService, crErr := NewLocalCryptoService()
	if crErr != nil {
		return nil, fmt.Errorf("failed to initialize crypto service: %w", crErr)
	}

	result.pool = newPool
	result.logger = logrus.New()
	result.logger.SetFormatter(&logrus.JSONFormatter{})

	result.blockchainClient = blockchainClient
	result.cryptoService = cryptoService

	return result, nil
}
