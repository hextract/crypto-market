package implementation

import (
	"context"

	"github.com/h4x4d/crypto-market/main/internal/models"
)

func (ds *DatabaseService) GetBidByID(ID string) (*models.Bid, error) {
	query := `
        SELECT 
            b.id,
            b.user_id, 
    		b.from_id, 
    		b.to_id,
    		b.status, 
    		b.create_date, 
    		b.complete_date, 
    		b.min_price, 
    		b.max_price, 
    		b.amount_to_buy, 
    		b.bought_amount, 
    		b.buy_speed, 
    		b.avg_price, 
        WHERE b.id = $1;`

	args := []any{ID}

	rows, err := ds.pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bid := &models.Bid{}
	var userID string
	err = rows.Scan(
		&bid.ID,
		&userID,
		&bid.FromCurrency,
		&bid.ToCurrency,
		&bid.Status,
		&bid.CreateDate,
		&bid.CompleteDate,
		&bid.MinPrice,
		&bid.MaxPrice,
		&bid.AmountToBuy,
		&bid.BoughtAmount,
		&bid.BuySpeed,
		&bid.AvgPrice,
	)
	if err != nil {
		return nil, err
	}

	return bid, nil
}
