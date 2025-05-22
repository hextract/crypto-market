package queue

import "context"

type IQueueConsumer interface {
	Consume(context.Context, Topics, Handler) error
	Close()
}
