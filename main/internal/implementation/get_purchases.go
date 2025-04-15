package implementation

import (
	"context"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"time"
)

func (ds *DatabaseService) GetPurchases(
	user *models.User,
	status *string,
	dateFrom, dateTo *time.Time,
	limit, offset *int64,
) ([]*models.Purchase, error) {
	query := `
        SELECT 
            b.id,
            cf.name AS currency_from,
            ct.name AS currency_to,
            COALESCE(b.bought_amount * b.avg_price, 0) AS amount_from,
            b.bought_amount AS amount_to,
            b.status,
            EXTRACT(EPOCH FROM b.create_date)::bigint AS date
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

	result := []*models.Purchase{}
	for rows.Next() {
		purchase := &models.Purchase{}
		err := rows.Scan(
			&purchase.ID,
			&purchase.CurrencyFrom,
			&purchase.CurrencyTo,
			&purchase.AmountFrom,
			&purchase.AmountTo,
			&purchase.Status,
			&purchase.Date,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, purchase)
	}

	return result, nil
}
