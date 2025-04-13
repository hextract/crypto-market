package models

import "time"

type Transfer struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Currency   string    `json:"currency"`
	Amount     float32   `json:"amount"`
	Commission float32   `json:"commission"`
	Operation  string    `json:"operation"`
	Status     string    `json:"status"`
	Date       time.Time `json:"date"`
	Address    string    `json:"address"`
}
