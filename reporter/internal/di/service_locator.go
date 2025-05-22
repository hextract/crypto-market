package di

import (
	"log"
	"reporter/internal/domain/services"
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
	OrderResultNotifier services.IOrderResultNotifier
}

func newServiceLocator(orderResultNotifier services.IOrderResultNotifier) (*ServiceLocator, error) {
	serviceLocator := &ServiceLocator{}
	serviceLocator.OrderResultNotifier = orderResultNotifier
	return serviceLocator, nil
}
