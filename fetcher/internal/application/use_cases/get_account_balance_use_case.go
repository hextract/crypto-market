package use_cases

import (
	"fetcher/internal/domain/entities"
	"fetcher/internal/domain/ports/repositories"
)

type GetAccountBalanceUseCase struct {
	accountBalanceRepository repositories.IAccountBalanceRepository
}

func (useCase *GetAccountBalanceUseCase) Execute() (*entities.AccountBalance, error) {
	return useCase.accountBalanceRepository.GetBalance()
}

func NewGetAccountBalanceUseCase(accountBalanceRepository repositories.IAccountBalanceRepository) (*GetAccountBalanceUseCase, error) {
	return &GetAccountBalanceUseCase{accountBalanceRepository: accountBalanceRepository}, nil
}
