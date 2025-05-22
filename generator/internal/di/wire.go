//go:build wireinject

package di

import (
	"generator/internal/application/use_cases"
	"generator/internal/di/providers"
	"generator/internal/domain/ports/queue"
	isenders "generator/internal/domain/ports/senders"
	"generator/internal/infrastructure/adapters/queues/kafka"
	"generator/internal/infrastructure/adapters/senders"
	"github.com/google/wire"
)

import (
	irepositories "generator/internal/domain/ports/repositories"
	iservices "generator/internal/domain/services"
	"generator/internal/infrastructure/adapters/repositories"
	"generator/internal/infrastructure/services"
)

func InitializeServiceLocator() (*ServiceLocator, error) {
	wire.Build(
		newServiceLocator,
		use_cases.NewGetSnapshotUseCase,
		use_cases.NewSendContinuousOrderUseCase,
		use_cases.NewCreateContinuousOrderUseCase,
		wire.NewSet(
			providers.SnapshotsTopicProvider,
			providers.KafkaBrokersProvider,
		),
		wire.NewSet(
			kafka.NewConsumer,
			wire.Bind(new(queue.IQueueConsumer), new(*kafka.Consumer)),
		),
		wire.NewSet(
			repositories.NewOrderBookSnaphotsRepository,
			wire.Bind(new(irepositories.IOrderBookSnaphotsRepository), new(*repositories.OrderBookSnaphotsRepository)),
		),
		wire.NewSet(
			services.NewSyntheticContinuousOrdersProcessor,
			wire.Bind(new(iservices.ISyntheticContinuousOrdersProcessor), new(*services.SyntheticContinuousOrdersProcessor)),
		),
		wire.NewSet(
			senders.NewContinuousOrderSender,
			wire.Bind(new(isenders.IContinuousOrderSender), new(*senders.ContinuousOrderSender)),
		),
	)
	return nil, nil
}
