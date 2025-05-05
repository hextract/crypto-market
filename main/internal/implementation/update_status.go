package implementation

import (
	"context"
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/jackc/pgx/v5"
)

func (ds *DatabaseService) UpdateOrderStatus(orderID string, status string, boughtAmount *float32) (string, error) {
	if status != "cancelled" && status != "finished" {
		return "", fmt.Errorf("invalid status: %s", status)
	}

	var completeDate strfmt.DateTime
	var setClause string
	args := []interface{}{orderID, status}

	if status == "cancelled" {
		setClause = "status = $2, complete_date = NULL, bought_amount = NULL"
	} else {
		completeDate = strfmt.DateTime(time.Now())
		setClause = "status = $2, complete_date = $3"
		args = append(args, completeDate)
		if boughtAmount != nil {
			setClause += ", bought_amount = $4"
			args = append(args, *boughtAmount)
		} else {
			return "", fmt.Errorf("invalid boughtAmount: %s", boughtAmount)
		}
	}

	query := fmt.Sprintf(`
		UPDATE bids 
		SET %s
		WHERE id = $1
		RETURNING id;
	`, setClause)

	var updatedID string
	err := ds.pool.QueryRow(context.Background(), query, args...).Scan(&updatedID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("order with id %s not found", orderID)
		}
		return "", fmt.Errorf("failed to update order status: %w", err)
	}

	return updatedID, nil
}
