package kafka

import (
	"context"
	"fetcher/internal/domain/ports/queue"
	"fetcher/internal/infrastructure/adapters/queues/kafka/utils"
	"github.com/google/uuid"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
	"log"
)

// do not edit this line
var consumerInterfaceSatisfiedConstraint queue.IQueueConsumer = (*Consumer)(nil)

type Consumer struct {
	admin   *kadm.Client
	clients []*kgo.Client
	brokers []string
}

func (consumer *Consumer) Consume(context context.Context, topics queue.Topics, handler queue.Handler) error {
	for _, topic := range topics {
		errCreate := utils.CreateTopic(consumer.admin, topic)
		if errCreate != nil {
			return errCreate
		}
	}

	groupID := uuid.New().String()
	client, errClient := kgo.NewClient(
		kgo.SeedBrokers(consumer.brokers...),
		kgo.ConsumerGroup(groupID),
		kgo.ConsumeTopics(topics...),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	)
	if errClient != nil {
		return errClient
	}
	consumer.clients = append(consumer.clients, client)

	go func() {
		errConsume := consumer.consume(context, handler, len(consumer.clients)-1)
		if errConsume != nil {
			panic("Failed to consume kafka")
		}
	}()
	return nil
}

func (consumer *Consumer) consume(context context.Context, handler queue.Handler, clientInd int) error {
	client := consumer.clients[clientInd]
	for {
		fetches := client.PollFetches(context)
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			errCallback := handler(record.Value)
			if errCallback != nil {
				log.Println("Error during handling kafka message", errCallback)
				continue
			}
			//log.Println(fmt.Sprintf("Received kafka message from topic '%s'"), record.Topic)
		}
	}

}

func (consumer *Consumer) Close() {
	consumer.admin.Close()
	for _, client := range consumer.clients {
		client.Close()
	}
}

func NewConsumer(brokers []string) (*Consumer, error) {
	consumer := &Consumer{}
	client, errClient := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if errClient != nil {
		return nil, errClient
	}

	admin := kadm.NewClient(client)
	consumer.admin = admin
	consumer.brokers = brokers
	return consumer, nil
}
