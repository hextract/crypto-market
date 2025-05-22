package services

import "executor/internal/domain/entities"

type IOrderExecutor interface {
	Execute(*entities.Order) (*entities.OrderResult, error)
}
