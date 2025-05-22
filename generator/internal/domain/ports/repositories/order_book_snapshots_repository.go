package repositories

import "generator/internal/domain/entities"

type IOrderBookSnaphotsRepository interface {
	Get() (*entities.OrderBookSnapshot, error)
}
