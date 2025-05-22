package repositories

import "fetcher/internal/domain/entities"

type IOrderBookRepository interface {
	GetData() (*entities.OrderBookData, error)
	ReceiveSnapshot(*entities.OrderBookSnapshot) error
}
