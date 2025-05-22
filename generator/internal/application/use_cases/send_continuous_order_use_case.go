package use_cases

import (
	"generator/internal/domain/entities"
	"generator/internal/domain/ports/senders"
)

type SendContinuousOrderUseCase struct {
	sender senders.IContinuousOrderSender
}

func NewSendContinuousOrderUseCase(sender senders.IContinuousOrderSender) (*SendContinuousOrderUseCase, error) {
	return &SendContinuousOrderUseCase{sender: sender}, nil
}

func (useCase *SendContinuousOrderUseCase) Execute(order *entities.ContinuousOrder) error {
	return useCase.sender.Send(order)
}
