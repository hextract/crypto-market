package entrypoints

import (
	"context"
	"executor/internal/di"
	"log"
)

func OrderProcessorEntryPoint(context context.Context) {
	errOrderProcessor := di.GetServiceLocator().OrderProcessor.Run(context)
	if errOrderProcessor != nil {
		log.Fatal("Failed to run order processor: ", errOrderProcessor)
	}
	log.Println("Order processor is started up")
}
