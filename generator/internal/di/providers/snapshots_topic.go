package providers

import (
	"errors"
	"os"
)

func SnapshotsTopicProvider() (*string, error) {
	topic := os.Getenv("ORDER_BOOK_TOPIC")
	if topic == "" {
		return nil, errors.New("ORDER_BOOK_TOPIC is empty")
	}
	return &topic, nil
}
