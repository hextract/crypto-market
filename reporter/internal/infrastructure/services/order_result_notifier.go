package services

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"reporter/internal/domain/entities"
	"reporter/internal/domain/ports/delivery"
	"reporter/internal/domain/ports/queue"
	"reporter/internal/domain/services"
)

// do not edit this line
var orderResultNotifierInterfaceSatisfiedConstraint services.IOrderResultNotifier = (*OrderResultNotifier)(nil)

type OrderResultNotifier struct {
	queueConsumer     queue.IQueueConsumer
	orderResultSender delivery.IOrderResultSender
}

func (notifier *OrderResultNotifier) Run(context.Context) error {
	resultsTopic := os.Getenv("RESULTS_TOPIC")
	if resultsTopic == "" {
		return errors.New("RESULTS_TOPIC environment variable not set")
	}
	errConsume := notifier.queueConsumer.Consume(context.Background(), []string{resultsTopic}, func(message queue.Message) error {
		orderResult := &entities.OrderResult{}
		errJson := json.Unmarshal(message, orderResult)
		if errJson != nil {
			return errJson
		}
		errorNotify := notifier.Notify(orderResult)
		if errorNotify != nil {
			return errorNotify
		}
		return nil
	})
	if errConsume != nil {
		return errConsume
	}
	return nil
}

func (notifier *OrderResultNotifier) Notify(orderResult *entities.OrderResult) error {
	return notifier.orderResultSender.Send(orderResult)
}

func NewOrderResultNotifier(queueConsumer queue.IQueueConsumer, orderResultSender delivery.IOrderResultSender) (*OrderResultNotifier, error) {
	orderResultNotifier := &OrderResultNotifier{}
	orderResultNotifier.queueConsumer = queueConsumer
	orderResultNotifier.orderResultSender = orderResultSender
	return orderResultNotifier, nil
}
