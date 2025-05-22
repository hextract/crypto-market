//go:build wireinject

package di

import (
	"executor/internal/di/providers"
	"executor/internal/domain/ports/queue"
	iservices "executor/internal/domain/services"
	"executor/internal/infrastructure/adapters/queues/kafka"
	"executor/internal/infrastructure/services"
	"github.com/google/wire"
)

func InitializeServiceLocator() (*ServiceLocator, error) {
	wire.Build(
		newServiceLocator,
		providers.KafkaBrokersProvider,
		wire.NewSet(
			services.NewOrderProcessor,
			wire.Bind(new(iservices.IOrderProcessor), new(*services.OrderProcessor)),
		),
		wire.NewSet(
			kafka.NewConsumer,
			wire.Bind(new(queue.IQueueConsumer), new(*kafka.Consumer)),
		),
		wire.NewSet(
			kafka.NewProducer,
			wire.Bind(new(queue.IQueueProducer), new(*kafka.Producer)),
		),
		wire.NewSet(
			services.NewOrderExecutor,
			wire.Bind(new(iservices.IOrderExecutor), new(*services.OrderExecutor)),
		),
	)
	return nil, nil
}
