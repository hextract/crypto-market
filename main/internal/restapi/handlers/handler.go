package handlers

import "github.com/h4x4d/crypto-market/main/internal/implementation"

type Handler struct {
	Database *implementation.DatabaseService
	MatchingEngine implementation.MatchingEngineService
}

func NewHandler(connStr string, blockchainClient *implementation.BlockchainClient) (*Handler, error) {
	db, err := implementation.NewDatabaseService(connStr, blockchainClient)
	if err != nil {
		return nil, err
	}

	return &Handler{Database: db, 
		MatchingEngine: implementation.NewMatcingEngineStub(), 
		}, nil
}
