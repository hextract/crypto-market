package entrypoints

import (
	"context"
	"fetcher/internal/di"
	"log"
)

func BinancePollerEntryPoint(context context.Context) {
	serviceLocator := di.GetServiceLocator()
	log.Println("BinancePoller is started up")
	errPoll := serviceLocator.ExchangePoller.PollOrderBook(context)
	if errPoll != nil {
		log.Fatal("Error during starting")
	}
	log.Println("BinancePoller is started up")
}
