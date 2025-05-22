package use_cases

import (
	"reporter/internal/domain/entities"
	"reporter/internal/domain/ports"
)

type SendOrderResultUseCase struct {
	orderResultSender ports.IOrderResultSender
}

func (useCase *SendOrderResultUseCase) Execute(order *entities.OrderResult) error {
	return useCase.orderResultSender.Send(order)
}
