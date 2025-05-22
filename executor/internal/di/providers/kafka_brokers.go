package providers

import (
	"errors"
	"fmt"
	"os"
)

func KafkaBrokersProvider() ([]string, error) {
	brokerHost := "redpanda"
	brokerPort, portExists := os.LookupEnv("REDPANDA_REST_PORT")
	if !portExists {
		return nil, errors.New("REDPANDA_REST_PORT environment variable not set")
	}

	brokers := []string{fmt.Sprintf("%s:%s", brokerHost, brokerPort)}
	return brokers, nil
}
