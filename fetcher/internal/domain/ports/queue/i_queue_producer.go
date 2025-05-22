package queue

type IQueueProducer interface {
	Publish(Message, Topics) error
	Close()
}
