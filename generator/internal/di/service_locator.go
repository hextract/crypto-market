package di

import (
	"generator/internal/domain/services"
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
	SyntheticContinuousOrdersProcessor services.ISyntheticContinuousOrdersProcessor
}

func newServiceLocator(syntheticContinuousOrdersProcessor services.ISyntheticContinuousOrdersProcessor) (*ServiceLocator, error) {
	serviceLocator := &ServiceLocator{}
	serviceLocator.SyntheticContinuousOrdersProcessor = syntheticContinuousOrdersProcessor
	return serviceLocator, nil
}
