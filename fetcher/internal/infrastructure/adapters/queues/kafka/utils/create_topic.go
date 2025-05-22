package utils

import (
	"context"
	"github.com/twmb/franz-go/pkg/kadm"
	"log"
)

func CreateTopic(admin *kadm.Client, topic string) error {
	exists, errExists := topicExists(admin, topic)
	if errExists != nil {
		return errExists
	}
	if exists {
		return nil
	}

	responses, err := admin.CreateTopics(context.Background(), 1, 1, nil, topic)
	if err != nil {
		return err
	}

	for _, ctr := range responses {
		if ctr.Err != nil {
			log.Println("Unable to create topic '%s' with error: %s", ctr.Topic, ctr.Err)
		} else {
			log.Println("Created topic '%s'", ctr.Topic)
		}
	}
	return nil
}
