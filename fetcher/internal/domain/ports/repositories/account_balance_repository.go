package repositories

import "fetcher/internal/domain/entities"

type IAccountBalanceRepository interface {
	GetBalance() (*entities.AccountBalance, error)
}
