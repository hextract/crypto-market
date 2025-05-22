package use_cases

import (
	"errors"
	"generator/internal/domain/entities"
)

const (
	firstSyntheticContinuousOrderId int64   = 1000000000
	volumeCoefficient               float64 = 0.1
	secondsPerHour                          = 3600
)

type CreateContinuousOrderFromOrderBookSnapshotUseCase struct {
	currentSyntheticContinuousOrderId int64
}

func (useCase *CreateContinuousOrderFromOrderBookSnapshotUseCase) Execute(snapshot *entities.OrderBookSnapshot) (*entities.ContinuousOrder, error) {
	if snapshot == nil {
		return nil, errors.New("snapshot cannot be nil")
	}

	result, errOrder := entities.NewContinousOrder()
	if errOrder != nil {
		return nil, errOrder
	}

	result.OrderID = useCase.currentSyntheticContinuousOrderId
	useCase.currentSyntheticContinuousOrderId++
	result.IsBid = useCase.currentSyntheticContinuousOrderId%2 == 0

	if result.IsBid {
		result.MaxSpeed = volumeCoefficient * snapshot.BidVolume * secondsPerHour
		result.Amount = snapshot.BidVolume
	} else {
		result.MaxSpeed = volumeCoefficient * snapshot.AskVolume * secondsPerHour
		result.Amount = snapshot.AskVolume
	}

	result.PriceLow = snapshot.BestBid
	result.PriceHigh = snapshot.BestAsk
	if result.PriceLow > result.PriceHigh {
		result.PriceLow, result.PriceHigh = result.PriceHigh, result.PriceLow
	}
	return result, nil
}

func NewCreateContinuousOrderUseCase() (*CreateContinuousOrderFromOrderBookSnapshotUseCase, error) {
	return &CreateContinuousOrderFromOrderBookSnapshotUseCase{
		currentSyntheticContinuousOrderId: firstSyntheticContinuousOrderId,
	}, nil
}
