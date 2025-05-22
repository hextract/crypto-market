package senders

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"generator/internal/domain/entities"
	"generator/internal/domain/ports/senders"
	"generator/internal/infrastructure/mappers"
	"net/http"
	"os"
)

var continuousOrderSenderInterfaceSatisfiedConstraint senders.IContinuousOrderSender = (*ContinuousOrderSender)(nil)

type ContinuousOrderSender struct{}

func (sender *ContinuousOrderSender) Send(order *entities.ContinuousOrder) error {
	dto, err := mappers.ContinuousOrderToJsonDto(order)
	if err != nil {
		return err
	}

	body, _ := json.Marshal(dto)

	matchingEngineHost := os.Getenv("MATCHING_ENGINE_HOST")
	matchingEnginePort := os.Getenv("MATCHING_ENGINE_PORT")
	if matchingEngineHost == "" || matchingEnginePort == "" {
		return errors.New("MATCHING_ENGINE_HOST or MATCHING_ENGINE_PORT environment variables not set")
	}
	req, _ := http.NewRequest("POST",
		fmt.Sprintf("http://%s:%s/create-order", matchingEngineHost, matchingEnginePort),
		bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil // или return errors.New("not implemented") если хочешь ошибку
}

func NewContinuousOrderSender() (*ContinuousOrderSender, error) {
	return &ContinuousOrderSender{}, nil
}
