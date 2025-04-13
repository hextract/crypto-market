package models

type User struct {
	UserID string `json:"user_id"`
}

type UserCurrency struct {
	Currency string  `json:"currency"`
	Amount   float32 `json:"amount"`
}
