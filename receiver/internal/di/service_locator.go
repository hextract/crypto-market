package di

import (
	"log"
	"receiver/internal/domain/ports/repositories"
	"receiver/internal/presentation/web/controllers"
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
			log.Fatal("Error during initializing ServiceLocator")
		}
	})
	return instance
}

type ServiceLocator struct {
	OrdersRepository repositories.IOrdersRepository
	OrdersController *controllers.OrdersController
}

func newServiceLocator(ordersRepository repositories.IOrdersRepository, ordersController *controllers.OrdersController) (*ServiceLocator, error) {
	serviceLocator := &ServiceLocator{}
	serviceLocator.OrdersRepository = ordersRepository
	serviceLocator.OrdersController = ordersController
	return serviceLocator, nil
}
