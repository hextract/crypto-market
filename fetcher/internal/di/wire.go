//go:build wireinject

package di

import (
	"fetcher/internal/application/use_cases"
	"fetcher/internal/di/providers"
	"fetcher/internal/domain/ports/queue"
	irepositories "fetcher/internal/domain/ports/repositories"
	iservices "fetcher/internal/domain/services"
	"fetcher/internal/infrastructure/adapters/queues/kafka"
	"fetcher/internal/infrastructure/adapters/repositories"
	"fetcher/internal/infrastructure/services"
	"fetcher/internal/presentation/web/controllers"
	"github.com/google/wire"
)

func InitializeServiceLocator() (*ServiceLocator, error) {
	wire.Build(
		controllers.NewMarketDataController,
		controllers.NewAccountBalanceController,
		newServiceLocator,
		providers.KafkaBrokersProvider,
		use_cases.NewGetOrderBookDataUseCase,
		use_cases.NewGetAccountBalanceUseCase,
		wire.NewSet(
			repositories.NewOrderBookRepository,
			wire.Bind(new(irepositories.IOrderBookRepository), new(*repositories.OrderBookRepository)),
		),
		wire.NewSet(
			repositories.NewAccountBalanceRepository,
			wire.Bind(new(irepositories.IAccountBalanceRepository), new(*repositories.AccountBalanceRepository)),
		),
		wire.NewSet(
			kafka.NewProducer,
			wire.Bind(new(queue.IQueueProducer), new(*kafka.Producer)),
		),
		wire.NewSet(
			kafka.NewConsumer,
			wire.Bind(new(queue.IQueueConsumer), new(*kafka.Consumer)),
		),
		wire.NewSet(
			services.NewBinancePoller,
			wire.Bind(new(iservices.IExchangePoller), new(*services.BinancePoller)),
		),
	)
	return nil, nil
}
