package services

import (
	"context"
	"executor/internal/domain/entities"
)

type IOrderProcessor interface {
	Process(*entities.Order) error
	Run(context.Context) error
}
