// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"reporter/internal/di/providers"
	"reporter/internal/infrastructure/adapters/delivery"
	"reporter/internal/infrastructure/adapters/queues/kafka"
	"reporter/internal/infrastructure/services"
)

// Injectors from wire.go:

func InitializeServiceLocator() (*ServiceLocator, error) {
	v, err := providers.KafkaBrokersProvider()
	if err != nil {
		return nil, err
	}
	consumer, err := kafka.NewConsumer(v)
	if err != nil {
		return nil, err
	}
	matchingEngineConfig := providers.MatchingEngineConfigProvider()
	clientService := providers.ClientServiceProvider(matchingEngineConfig)
	orderResultSender, err := delivery.NewOrderResultSender(clientService)
	if err != nil {
		return nil, err
	}
	orderResultNotifier, err := services.NewOrderResultNotifier(consumer, orderResultSender)
	if err != nil {
		return nil, err
	}
	serviceLocator, err := newServiceLocator(orderResultNotifier)
	if err != nil {
		return nil, err
	}
	return serviceLocator, nil
}
