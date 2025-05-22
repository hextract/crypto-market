package kafka

import (
	"context"
	"fetcher/internal/domain/ports/queue"
	"fetcher/internal/infrastructure/adapters/queues/kafka/utils"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

// do not edit this line
var publisherInterfaceSatisfiedConstraint queue.IQueueProducer = (*Producer)(nil)

type Producer struct {
	admin  *kadm.Client
	client *kgo.Client
}

func (producer *Producer) Publish(message queue.Message, topics queue.Topics) error {
	for _, topic := range topics {
		errCreate := utils.CreateTopic(producer.admin, topic)
		if errCreate != nil {
			return errCreate
		}

		record := &kgo.Record{Topic: topic, Value: message}
		producer.client.Produce(context.Background(), record, nil)
	}
	return nil
}

func (producer *Producer) Close() {
	producer.admin.Close()
	producer.client.Close()
}

func NewProducer(brokers []string) (*Producer, error) {
	producer := &Producer{}
	client, errClient := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if errClient != nil {
		return nil, errClient
	}
	producer.client = client

	admin := kadm.NewClient(client)
	producer.admin = admin
	return producer, nil
}
