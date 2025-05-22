package use_cases

import (
	"receiver/internal/domain/entities"
	"receiver/internal/domain/ports/repositories"
)

type ReceiveOrderUseCase struct {
	ordersRepository repositories.IOrdersRepository
}

func (useCase *ReceiveOrderUseCase) Execute(order *entities.Order) error {
	return useCase.ordersRepository.Store(order)
}

func NewReceiveOrderUseCase(ordersRepository repositories.IOrdersRepository) (*ReceiveOrderUseCase, error) {
	return &ReceiveOrderUseCase{ordersRepository: ordersRepository}, nil
}
