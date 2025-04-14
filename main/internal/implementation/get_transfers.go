package implementation

import (
	"context"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"time"
)

func (ds *DatabaseService) GetTransfers(
	user *models.User,
	minAmount, maxAmount *float32,
	status, currency, operation *string,
	dateFrom, dateTo *time.Time,
	limit, offset *int64,
) ([]*models.Transfer, error) {
	query := `
        SELECT 
            t.id,
            c.name AS currency,
            t.amount,
            t.commission,
            t.type AS operation,
            t.status,
            EXTRACT(EPOCH FROM t.created_at)::bigint AS date,
            t.address
        FROM transactions t
        JOIN currencies c ON t.currency_id = c.currency_id
        WHERE t.user_id = $1`

	args := []interface{}{user.UserID}
	argIndex := 2

	if minAmount != nil {
		query += " AND t.amount >= $" + string(rune('0'+argIndex))
		args = append(args, *minAmount)
		argIndex++
	}
	if maxAmount != nil {
		query += " AND t.amount <= $" + string(rune('0'+argIndex))
		args = append(args, *maxAmount)
		argIndex++
	}
	if status != nil {
		query += " AND t.status = $" + string(rune('0'+argIndex))
		args = append(args, *status)
		argIndex++
	}
	if currency != nil {
		query += " AND c.name = $" + string(rune('0'+argIndex))
		args = append(args, *currency)
		argIndex++
	}
	if operation != nil {
		query += " AND t.type = $" + string(rune('0'+argIndex))
		args = append(args, *operation)
		argIndex++
	}
	if dateFrom != nil {
		query += " AND t.created_at >= $" + string(rune('0'+argIndex))
		args = append(args, *dateFrom)
		argIndex++
	}
	if dateTo != nil {
		query += " AND t.created_at <= $" + string(rune('0'+argIndex))
		args = append(args, *dateTo)
		argIndex++
	}

	// Добавляем ORDER BY перед LIMIT и OFFSET
	query += " ORDER BY t.created_at DESC"

	// Добавляем LIMIT и OFFSET, если указаны
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

	// Точка с запятой в конце
	query += ";"

	rows, err := ds.pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*models.Transfer
	for rows.Next() {
		transfer := &models.Transfer{}
		err := rows.Scan(
			&transfer.ID,
			&transfer.Currency,
			&transfer.Amount,
			&transfer.Commission,
			&transfer.Operation,
			&transfer.Status,
			&transfer.Date,
			&transfer.Address,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, transfer)
	}

	return result, nil
}
