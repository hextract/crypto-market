package implementation

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"time"
)

func (ds *DatabaseService) CreateBid(userID string, fromCurrency, toCurrency string, minPrice, maxPrice, amountToBuy, buySpeed float32) (string, error) {
	bidID := fmt.Sprintf("bid_%d", int(time.Now().UnixNano()))

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

func (ds *DatabaseService) CheckBid(bid int64) (string, error) {
	bidID := fmt.Sprintf("bid_%d", int(bid))

	query := `select id from bids where id=$1;`

	var foundID string
	err := ds.pool.QueryRow(context.Background(), query, bidID).Scan(&foundID)

	switch {
	case err == nil:
		return foundID, nil
	case errors.Is(err, pgx.ErrNoRows):
		return "", fmt.Errorf("bid with id %s not found", bidID)
	default:
		return "", fmt.Errorf("failed to check bid existence: %w", err)
	}
}
