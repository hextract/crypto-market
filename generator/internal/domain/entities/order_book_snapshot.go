package entities

type OrderBookSnapshot struct {
	Timestamp  int64   `json:"timestamp"`
	AskVolume  float64 `json:"ask_volume"`
	BidVolume  float64 `json:"bid_volume"`
	Depth      int32   `json:"depth"`
	Mid        float64 `json:"mid"`
	Spread     float64 `json:"spread"`
	BestBid    float64 `json:"best_bid"`
	BestAsk    float64 `json:"best_ask"`
	VWAPBid    float64 `json:"vwap_bid"`
	VWAPAsk    float64 `json:"vwap_ask"`
	Imbalance  float64 `json:"imbalance"`
	LiqBid1Pct float64 `json:"liq_bid_1%"`
	LiqAsk1Pct float64 `json:"liq_ask_1%"`
}
