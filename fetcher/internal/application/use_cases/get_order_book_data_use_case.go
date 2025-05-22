package use_cases

import (
	"fetcher/internal/domain/entities"
	"fetcher/internal/domain/ports/repositories"
)

type GetOrderBookDataUseCase struct {
	orderBookRepository repositories.IOrderBookRepository
}

func (useCase *GetOrderBookDataUseCase) Execute() (*entities.OrderBookData, error) {
	return useCase.orderBookRepository.GetData()
}

func NewGetOrderBookDataUseCase(orderBookRepository repositories.IOrderBookRepository) (*GetOrderBookDataUseCase, error) {
	return &GetOrderBookDataUseCase{orderBookRepository: orderBookRepository}, nil
}
