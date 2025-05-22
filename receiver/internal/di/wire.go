//go:build wireinject

package di

import (
	"github.com/google/wire"
	"receiver/internal/di/providers"
	"receiver/internal/domain/ports/queue"
	irepositories "receiver/internal/domain/ports/repositories"
	"receiver/internal/infrastructure/adapters/queues/kafka"
	"receiver/internal/infrastructure/adapters/repositories"
	"receiver/internal/presentation/web/controllers"
)

func InitializeServiceLocator() (*ServiceLocator, error) {
	wire.Build(
		newServiceLocator,
		controllers.NewOrdersController,
		wire.NewSet(
			repositories.NewOrdersRepository,
			wire.Bind(new(irepositories.IOrdersRepository), new(*repositories.OrdersRepository)),
		),
		wire.NewSet(
			kafka.NewProducer,
			wire.Bind(new(queue.IQueueProducer), new(*kafka.Producer)),
		),
		providers.KafkaBrokersProvider,
	)
	return nil, nil
}
