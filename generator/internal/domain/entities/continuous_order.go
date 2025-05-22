package entities

type ContinuousOrder struct {
	Pair      TradingPair
	OrderID   int64
	IsBid     bool
	MaxSpeed  float64
	Amount    float64
	PriceLow  float64
	PriceHigh float64
}

func NewContinousOrder() (*ContinuousOrder, error) {
	order := &ContinuousOrder{}
	order.Pair = TradingPair{Base: "ETH", Quote: "USDT"}
	return order, nil
}
