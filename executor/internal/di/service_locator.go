package di

import (
	"executor/internal/domain/services"
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
			log.Fatal("Error during initializing ServiceLocator")
		}
	})
	return instance
}

type ServiceLocator struct {
	OrderProcessor services.IOrderProcessor
}

func newServiceLocator(orderProcessor services.IOrderProcessor) (*ServiceLocator, error) {
	serviceLocator := &ServiceLocator{}
	serviceLocator.OrderProcessor = orderProcessor
	return serviceLocator, nil
}
