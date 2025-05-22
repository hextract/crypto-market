package repositories

import (
	"encoding/json"
	"errors"
	"os"
	"receiver/internal/domain/entities"
	"receiver/internal/domain/ports/queue"
	"receiver/internal/domain/ports/repositories"
	"receiver/internal/infrastructure/mappers"
)

// do not edit this line
var interfaceSatisfiedConstraint repositories.IOrdersRepository = (*OrdersRepository)(nil)

type OrdersRepository struct {
	queueProducer queue.IQueueProducer
	ordersTopic   string
}

func (repository *OrdersRepository) Store(order *entities.Order) error {
	if order == nil {
		return errors.New("Order is nil")
	}

	jsonOrder, errJson := json.Marshal(mappers.OrderToJsonDto(order))
	if errJson != nil {
		return errJson
	}
	errPublish := repository.queueProducer.Publish(jsonOrder, []string{repository.ordersTopic})
	if errPublish != nil {
		return errPublish
	}
	return nil
}

func NewOrdersRepository(queueProducer queue.IQueueProducer) (*OrdersRepository, error) {
	repository := &OrdersRepository{}
	repository.queueProducer = queueProducer

	repository.ordersTopic = os.Getenv("ORDERS_TOPIC")
	if repository.ordersTopic == "" {
		return nil, errors.New("ORDERS_TOPIC is empty")
	}

	return repository, nil
}
