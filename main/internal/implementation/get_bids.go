package implementation

import (
	"context"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"time"
)

func (ds *DatabaseService) GetBids(
	user *models.User,
	status *string,
	dateFrom, dateTo *time.Time,
	limit, offset *int64,
) ([]*models.Bid, error) {
	query := `
        SELECT 
            b.id,
            cf.name AS currency_from,
            ct.name AS currency_to,
            b.status,
            b.create_date,
            b.complete_date,
            b.min_price,
            b.max_price,
            b.amount_to_buy,
            b.bought_amount,
            b.buy_speed,
            b.avg_price
        FROM bids b
        JOIN currencies cf ON b.from_id = cf.currency_id
        JOIN currencies ct ON b.to_id = ct.currency_id
        WHERE b.user_id = $1`

	args := []interface{}{user.UserID}
	argIndex := 2

	if status != nil {
		query += " AND b.status = $" + string(rune('0'+argIndex))
		args = append(args, *status)
		argIndex++
	}
	if dateFrom != nil {
		query += " AND b.create_date >= $" + string(rune('0'+argIndex))
		args = append(args, *dateFrom)
		argIndex++
	}
	if dateTo != nil {
		query += " AND b.create_date <= $" + string(rune('0'+argIndex))
		args = append(args, *dateTo)
		argIndex++
	}

	query += " ORDER BY b.create_date DESC"

	if limit != nil {
		query += " LIMIT $" + string(rune('0'+argIndex))
		args = append(args, *limit)
		argIndex++
	}
	if offset != nil {
		query += " OFFSET $" + string(rune('0'+argIndex))
		args = append(args, *offset)
		argIndex++
	}

	query += ";"

	rows, err := ds.pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*models.Bid
	for rows.Next() {
		bid := &models.Bid{}
		err := rows.Scan(
			&bid.ID,
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
		result = append(result, bid)
	}

	return result, nil
}
