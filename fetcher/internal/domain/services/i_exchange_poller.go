package services

import "context"

type IExchangePoller interface {
	PollOrderBook(context.Context) error
}
