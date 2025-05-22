package repositories

import (
	"container/list"
	"context"
	"encoding/json"
	"errors"
	"fetcher/internal/domain/entities"
	"fetcher/internal/domain/ports/queue"
	"fetcher/internal/domain/ports/repositories"
	"fetcher/internal/infrastructure/dto"
	"fetcher/internal/infrastructure/mappers"
	"os"
	"strconv"
	"sync"
)

// do not edit this line
var interfaceSatisfiedConstraint repositories.IOrderBookRepository = (*OrderBookRepository)(nil)

type OrderBookRepository struct {
	queueConsumer  queue.IQueueConsumer
	queueProducer  queue.IQueueProducer
	orderBookTopic string

	snapshotsLimit int
	snapshots      list.List
	mutex          sync.Mutex
}

func (repository *OrderBookRepository) GetData() (*entities.OrderBookData, error) {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()
	data := &entities.OrderBookData{}
	data.Snapshots = make([]*entities.OrderBookSnapshot, 0)
	for snapshot := repository.snapshots.Front(); snapshot != nil; snapshot = snapshot.Next() {
		data.Snapshots = append(data.Snapshots, snapshot.Value.(*entities.OrderBookSnapshot))
	}
	return data, nil
}

func (repository *OrderBookRepository) ReceiveSnapshot(snapshot *entities.OrderBookSnapshot) error {
	if snapshot == nil {
		return errors.New("OrderBookSnapshot is nil")
	}

	jsonSnapshot, errJson := json.Marshal(mappers.OrderBookSnapshotToJsonDto(snapshot))
	if errJson != nil {
		return errJson
	}
	errPublish := repository.queueProducer.Publish(jsonSnapshot, []string{repository.orderBookTopic})
	if errPublish != nil {
		return errPublish
	}
	return nil
}

func (repository *OrderBookRepository) consumeSnapshots() error {
	handler := func(message queue.Message) error {
		repository.mutex.Lock()
		defer repository.mutex.Unlock()

		jsonOrderBookSnapshot := &dto.JsonOrderBookSnapshot{}
		errUnmarshal := json.Unmarshal(message, jsonOrderBookSnapshot)
		if errUnmarshal != nil {
			return errUnmarshal
		}

		repository.snapshots.PushBack(mappers.OrderBookSnapshotToEntity(jsonOrderBookSnapshot))
		if repository.snapshots.Len() > repository.snapshotsLimit {
			repository.snapshots.Remove(repository.snapshots.Front())
		}
		return nil
	}
	errConsume := repository.queueConsumer.Consume(context.Background(), []string{repository.orderBookTopic}, handler)
	if errConsume != nil {
		return errConsume
	}
	return nil
}

func NewOrderBookRepository(queueConsumer queue.IQueueConsumer, queueProducer queue.IQueueProducer) (*OrderBookRepository, error) {
	repository := &OrderBookRepository{}
	repository.queueConsumer = queueConsumer
	repository.queueProducer = queueProducer

	repository.orderBookTopic = os.Getenv("ORDER_BOOK_TOPIC")
	if repository.orderBookTopic == "" {
		return nil, errors.New("ORDER_BOOK_TOPIC is empty")
	}

	snapshotsLimitString := os.Getenv("ORDER_BOOK_DATA_LIMIT")
	if snapshotsLimitString == "" {
		return nil, errors.New("ORDER_BOOK_DATA_LIMIT is empty")
	}
	snapshotLimit, errCastToInt := strconv.Atoi(snapshotsLimitString)
	if errCastToInt != nil {
		return nil, errCastToInt
	}
	repository.snapshotsLimit = snapshotLimit

	errConsumeSnapshots := repository.consumeSnapshots()
	if errConsumeSnapshots != nil {
		return nil, errConsumeSnapshots
	}
	return repository, nil
}
