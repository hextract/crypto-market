package use_cases

import (
	"fetcher/internal/domain/entities"
	"fetcher/internal/domain/ports/repositories"
)

type ReceiveOrderBookSnapshotUseCase struct {
	orderBookRepository repositories.IOrderBookRepository
}

func (useCase *ReceiveOrderBookSnapshotUseCase) Execute(orderBookSnapshot *entities.OrderBookSnapshot) error {
	return useCase.orderBookRepository.ReceiveSnapshot(orderBookSnapshot)
}

func NewReceiveOrderBookSnapshotUseCase(orderBookRepository repositories.IOrderBookRepository) (*ReceiveOrderBookSnapshotUseCase, error) {
	return &ReceiveOrderBookSnapshotUseCase{orderBookRepository: orderBookRepository}, nil
}
