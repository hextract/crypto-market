package orders_executor

import "executor/internal/binance"

type OrdersExecutor struct {
	Client *binance.BinanceClient
}

func New() (*OrdersExecutor, error) {
	client := binance.NewBinanceClient()
	return &OrdersExecutor{Client: client}, nil
}
