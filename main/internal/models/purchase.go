// models/models.go
package models

import "time"

type Purchase struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	CurrencyFrom string    `json:"currency_from"`
	CurrencyTo   string    `json:"currency_to"`
	AmountFrom   float32   `json:"amount_from"`
	AmountTo     float32   `json:"amount_to"`
	Status       string    `json:"status"`
	Date         time.Time `json:"date"`
}
