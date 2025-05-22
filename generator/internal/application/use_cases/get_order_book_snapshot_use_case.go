package use_cases

import (
	"generator/internal/domain/entities"
	"generator/internal/domain/ports/repositories"
)

type GetSnapshotUseCase struct {
	orderBookSnapshotsRepository repositories.IOrderBookSnaphotsRepository
}

func NewGetSnapshotUseCase(orderBookSnapshotsRepository repositories.IOrderBookSnaphotsRepository) (*GetSnapshotUseCase, error) {
	return &GetSnapshotUseCase{orderBookSnapshotsRepository: orderBookSnapshotsRepository}, nil
}

func (useCase *GetSnapshotUseCase) Execute() (*entities.OrderBookSnapshot, error) {
	return useCase.orderBookSnapshotsRepository.Get()
}
