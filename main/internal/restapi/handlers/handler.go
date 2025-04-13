package handlers

import "github.com/h4x4d/crypto-market/main/internal/implementation"

type Handler struct {
	Database *implementation.DatabaseService
}

func NewHandler(connStr string) (*Handler, error) {
	db, err := implementation.NewDatabaseService(connStr)
	if err != nil {
		return nil, err
	}

	return &Handler{db}, nil
}
