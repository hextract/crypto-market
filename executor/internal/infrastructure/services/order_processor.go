package services

import (
	"context"
	"encoding/json"
	"errors"
	"executor/internal/domain/entities"
	"executor/internal/domain/ports/queue"
	"executor/internal/domain/services"
	"executor/internal/infrastructure/mappers"
	"os"
)

// do not edit this line
var OrderProcessorInterfaceSatisfiedConstraint services.IOrderProcessor = (*OrderProcessor)(nil)

type OrderProcessor struct {
	queueConsumer     queue.IQueueConsumer
	queueProducer     queue.IQueueProducer
	executor          services.IOrderExecutor
	orderResultsTopic string
	ordersTopic       string
}

func (processor *OrderProcessor) Process(order *entities.Order) error {
	orderResult, errOrderResult := processor.executor.Execute(order)
	if errOrderResult != nil {
		return errOrderResult
	}
	jsonOrderResult, errJson := json.Marshal(mappers.OrderResultToJsonDto(orderResult))
	if errJson != nil {
		return errJson
	}
	return processor.queueProducer.Publish(jsonOrderResult, []string{processor.orderResultsTopic})
}

func (processor *OrderProcessor) Run(context context.Context) error {
	errConsume := processor.queueConsumer.Consume(context, []string{processor.ordersTopic}, func(message queue.Message) error {
		order := &entities.Order{}
		errJson := json.Unmarshal(message, order)
		if errJson != nil {
			return errJson
		}
		return processor.Process(order)
	})
	return errConsume
}

func NewOrderProcessor(queueConsumer queue.IQueueConsumer, queueProducer queue.IQueueProducer, executor services.IOrderExecutor) (*OrderProcessor, error) {
	processor := &OrderProcessor{}

	// incorrect topics should be passed from the outside
	processor.orderResultsTopic = os.Getenv("RESULTS_TOPIC")
	if processor.orderResultsTopic == "" {
		return nil, errors.New("RESULTS_TOPIC environment variable not set")
	}

	processor.ordersTopic = os.Getenv("ORDERS_TOPIC")
	if processor.ordersTopic == "" {
		return nil, errors.New("ORDERS_TOPIC environment variable not set")
	}

	processor.queueConsumer = queueConsumer
	processor.queueProducer = queueProducer
	processor.executor = executor
	return processor, nil
}
