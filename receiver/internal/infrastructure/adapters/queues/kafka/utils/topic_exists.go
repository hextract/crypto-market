package utils

import (
	"context"
	"github.com/twmb/franz-go/pkg/kadm"
)

func topicExists(admin *kadm.Client, topic string) (bool, error) {
	topicsMetadata, err := admin.ListTopics(context.Background())
	if err != nil {
		return false, err
	}
	for _, metadata := range topicsMetadata {
		if metadata.Topic == topic {
			return true, nil
		}
	}
	return false, nil
}
