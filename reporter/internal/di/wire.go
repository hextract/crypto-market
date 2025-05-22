//go:build wireinject

package di

import (
	"github.com/google/wire"
	"reporter/internal/di/providers"
	idelivery "reporter/internal/domain/ports/delivery"
	"reporter/internal/domain/ports/queue"
	iservices "reporter/internal/domain/services"
	"reporter/internal/infrastructure/adapters/delivery"
	"reporter/internal/infrastructure/adapters/queues/kafka"
	"reporter/internal/infrastructure/services"
)

func InitializeServiceLocator() (*ServiceLocator, error) {
	wire.Build(
		newServiceLocator,
		providers.KafkaBrokersProvider,
		wire.NewSet(
			services.NewOrderResultNotifier,
			wire.Bind(new(iservices.IOrderResultNotifier), new(*services.OrderResultNotifier)),
		),
		wire.NewSet(
			delivery.NewOrderResultSender,
			wire.Bind(new(idelivery.IOrderResultSender), new(*delivery.OrderResultSender)),
		),
		wire.NewSet(
			kafka.NewConsumer,
			wire.Bind(new(queue.IQueueConsumer), new(*kafka.Consumer)),
		),
		wire.NewSet(
			providers.MatchingEngineConfigProvider,
			providers.ClientServiceProvider,
		),
	)
	return nil, nil
}
