package use_cases

import (
	"executor/internal/domain/entities"
	"executor/internal/domain/services"
)

type ExecuteOrderUseCase struct {
	orderExecutor services.IOrderExecutor
}

func (useCase *ExecuteOrderUseCase) Execute(order *entities.Order) (*entities.OrderResult, error) {
	return useCase.orderExecutor.Execute(order)
}

func NewExecuteOrderUseCase(orderExecutor services.IOrderExecutor) (*ExecuteOrderUseCase, error) {
	useCase := &ExecuteOrderUseCase{}
	useCase.orderExecutor = orderExecutor
	return useCase, nil
}
