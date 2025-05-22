package dto

type JsonOrder struct {
	Indicator string  `json:"indicator"`
	Price     float64 `json:"price"`
	Ticker    string  `json:"ticker"`
	Type      string  `json:"type"`
	UserID    int64   `json:"user_id"`
	Volume    float64 `json:"volume"`
}
