package di

import (
	"fetcher/internal/domain/ports/repositories"
	"fetcher/internal/domain/services"
	"fetcher/internal/presentation/web/controllers"
	"log"
	"sync"
)

var (
	instance *ServiceLocator
	once     sync.Once
)

func GetServiceLocator() *ServiceLocator {
	once.Do(func() {
		var errInstance error
		instance, errInstance = InitializeServiceLocator()
		if errInstance != nil {
			log.Fatal("Error during initializing ServiceLocator " + errInstance.Error())
		}
	})
	return instance
}

type ServiceLocator struct {
	OrderBookRepository      repositories.IOrderBookRepository
	ExchangePoller           services.IExchangePoller
	MarketDataController     *controllers.MarketDataController
	AccountBalanceController *controllers.AccountBalanceController
}

func newServiceLocator(orderBookRepository repositories.IOrderBookRepository, exchangePoller services.IExchangePoller, marketDataController *controllers.MarketDataController, accountBalanceController *controllers.AccountBalanceController) (*ServiceLocator, error) {
	serviceLocator := &ServiceLocator{}
	serviceLocator.OrderBookRepository = orderBookRepository
	serviceLocator.ExchangePoller = exchangePoller
	serviceLocator.MarketDataController = marketDataController
	serviceLocator.AccountBalanceController = accountBalanceController
	return serviceLocator, nil
}
