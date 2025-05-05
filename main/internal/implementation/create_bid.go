package implementation

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (ds *DatabaseService) CreateBid(userID string, fromCurrency, toCurrency string, minPrice, maxPrice, amountToBuy, buySpeed float32) (string, error) {
	bidID := fmt.Sprintf("bid_%s", uuid.New().String())

	query := `
		INSERT INTO bids (
			id, 
			user_id, 
			from_id, 
			to_id, 
			min_price, 
			max_price, 
			amount_to_buy, 
			buy_speed,
			create_date
		) 
		SELECT 
			$1, $2, 
			(SELECT currency_id FROM currencies WHERE name = $3), 
			(SELECT currency_id FROM currencies WHERE name = $4), 
			$5, $6, $7, $8, $9
		RETURNING id;
	`

	args := []any{
		bidID,
		userID,
		fromCurrency,
		toCurrency,
		minPrice,
		maxPrice,
		amountToBuy,
		buySpeed,
		time.Now(),
	}

	var id string
	err := ds.pool.QueryRow(context.Background(), query, args...).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to create bid: %w", err)
	}

	return id, nil
}
