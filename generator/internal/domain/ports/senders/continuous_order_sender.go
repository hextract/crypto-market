package senders

import "generator/internal/domain/entities"

type IContinuousOrderSender interface {
	Send(order *entities.ContinuousOrder) error
}
