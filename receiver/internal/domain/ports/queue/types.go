package queue

type Message []byte
type Handler func(Message) error
type Topics []string
