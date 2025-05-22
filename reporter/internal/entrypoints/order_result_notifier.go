package entrypoints

import (
	"context"
	"log"
	"reporter/internal/di"
)

func EntryPoint(context context.Context) {
	errNotifier := di.GetServiceLocator().OrderResultNotifier.Run(context)
	if errNotifier != nil {
		log.Fatal("Failed to run order result notifier: ", errNotifier)
	}
	log.Println("OrderResultNotifier is started up")
}
