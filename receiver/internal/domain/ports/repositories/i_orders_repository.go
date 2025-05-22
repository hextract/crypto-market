package repositories

import "receiver/internal/domain/entities"

type IOrdersRepository interface {
	Store(*entities.Order) error
}
