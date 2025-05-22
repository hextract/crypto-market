package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"generator/internal/domain/entities"
	"generator/internal/domain/ports/queue"
	"generator/internal/domain/ports/repositories"
	"log"
	"sync"
)

var interfaceSatisfiedConstraint repositories.IOrderBookSnaphotsRepository = (*OrderBookSnaphotsRepository)(nil)

type OrderBookSnaphotsRepository struct {
	mu             sync.RWMutex
	lastSnapshot   *entities.OrderBookSnapshot
	queueConsumer  queue.IQueueConsumer
	snapshotsTopic string
}

func NewOrderBookSnaphotsRepository(
	queueConsumer queue.IQueueConsumer,
	snapshotsTopic *string,
) (*OrderBookSnaphotsRepository, error) {
	repo := &OrderBookSnaphotsRepository{
		queueConsumer:  queueConsumer,
		snapshotsTopic: *snapshotsTopic,
	}

	// запускаем потребление
	if err := repo.pollLastSnapshot(); err != nil {
		return nil, err
	}
	return repo, nil
}

func (repository *OrderBookSnaphotsRepository) Get() (*entities.OrderBookSnapshot, error) {
	repository.mu.RLock()
	defer repository.mu.RUnlock()

	if repository.lastSnapshot == nil {
		return nil, errors.New("no snapshot received yet")
	}
	// возвращаем копию, чтобы никто извне не менял внутренний объект
	copy := *repository.lastSnapshot
	return &copy, nil
}

func (repository *OrderBookSnaphotsRepository) pollLastSnapshot() error {
	return repository.queueConsumer.Consume(context.Background(),
		[]string{repository.snapshotsTopic},
		func(msg queue.Message) error {
			data := msg
			var snap entities.OrderBookSnapshot
			if err := json.Unmarshal(data, &snap); err != nil {
				log.Printf("OrderBookSnapshot unmarshal error: %v, raw: %s\n", err, data)
				return err
			}
			repository.mu.Lock()
			repository.lastSnapshot = &snap
			repository.mu.Unlock()

			return nil
		})
}

//func (s OrderSnapshotRepository) SendContinuousOrder(order *entities.ContinuousOrder) error {
//TODO implement me
//	panic("implement me")
//}

//func (s OrderSnapshotRepository) SendContinuousOrder(order *entities.ContinuousOrder) error {
//	dto, err := mappers.ContinuousOrderToJsonDto(order)
//	if err != nil {
//		return err
//	}
//
//	body, _ := json.Marshal(dto)
//
//	req, _ := http.NewRequest("POST", "http://localhost:8085/create-order", bytes.NewBuffer(body))
//	req.Header.Set("Content-Type", "application/json")
//
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//	return nil // или return errors.New("not implemented") если хочешь ошибку
//}
