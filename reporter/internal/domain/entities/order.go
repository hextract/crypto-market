package entities

type Order struct {
	Indicator string
	Price     float64
	Ticker    string
	Type      string
	UserID    int64
	Volume    float64
}
