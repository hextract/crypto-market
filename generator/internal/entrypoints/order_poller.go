package entrypoints

import (
	"context"
	"generator/internal/di"
	"log"
)

func SyntheticContinuousOrdersProcessorEntryPoint(context context.Context) {
	serviceLocator := di.GetServiceLocator()
	log.Println("OrderPoller is started up")
	errPoll := serviceLocator.SyntheticContinuousOrdersProcessor.StartProcessing(context)
	if errPoll != nil {
		log.Fatal("Error during starting")
	}
	log.Println("OrderPoller is started up")
}
