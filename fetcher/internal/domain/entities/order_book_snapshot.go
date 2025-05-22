package entities

type OrderBookSnapshot struct {
	Timestamp  int64
	AskVolume  float64
	BidVolume  float64
	Depth      int32
	Mid        float64
	Spread     float64
	BestBid    float64
	BestAsk    float64
	VWAPBid    float64
	VWAPAsk    float64
	Imbalance  float64
	LiqBid1Pct float64
	LiqAsk1Pct float64
}
