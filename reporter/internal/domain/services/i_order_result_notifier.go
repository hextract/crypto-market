package services

import (
	"context"
	"reporter/internal/domain/entities"
)

type IOrderResultNotifier interface {
	Notify(*entities.OrderResult) error
	Run(context.Context) error
}
