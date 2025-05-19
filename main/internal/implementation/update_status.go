package implementation

import (
	"context"
	"fmt"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/jackc/pgx/v5"
)

func (ds *DatabaseService) UpdateOrderStatus(
	orderID string,
	status string,
	boughtAmount *float32,
	paidPrice *float32,
) (prevBoughtAmount *float32, prevAvgPrice *float32, from_id int, err error) {
	tx, err := ds.pool.Begin(context.Background())
	if err != nil {
		return nil, nil, 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	err = tx.QueryRow(context.Background(),
		`SELECT bought_amount, avg_price, from_id FROM bids WHERE id = $1 FOR UPDATE`,
		orderID,
	).Scan(&prevBoughtAmount, &prevAvgPrice, &from_id)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil, 0, fmt.Errorf("order with id %s not found", orderID)
		}
		return nil, nil, 0, fmt.Errorf("failed to get current order values: %w", err)
	}

	var (
		setClauses []string
		args       []interface{}
		argPos     = 1
	)

	setClauses = append(setClauses, fmt.Sprintf("status = $%d", argPos))
	args = append(args, status)
	argPos++

	if status == "cancelled" {
		setClauses = append(setClauses, "complete_date = NULL")
	} else if status == "finished" {
		setClauses = append(setClauses, fmt.Sprintf("complete_date = $%d", argPos))
		args = append(args, strfmt.DateTime(time.Now()))
		argPos++
	}

	if boughtAmount != nil {
		setClauses = append(setClauses, fmt.Sprintf("bought_amount = $%d", argPos))
		args = append(args, *boughtAmount)
		argPos++

		if paidPrice != nil {
			setClauses = append(setClauses, fmt.Sprintf("avg_price = $%d", argPos))
			args = append(args, *paidPrice / *boughtAmount)
			argPos++
		}
	}

	query := fmt.Sprintf(
		`UPDATE bids SET %s WHERE id = $%d RETURNING id`,
		strings.Join(setClauses, ", "),
		argPos,
	)
	args = append(args, orderID)

	rows, err := tx.Query(context.Background(), query, args...)
	if err != nil {
		return prevBoughtAmount, prevAvgPrice, 0,
			fmt.Errorf("failed to update order: %w", err)
	}
	rows.Close()

	if err = tx.Commit(context.Background()); err != nil {
		return prevBoughtAmount, prevAvgPrice, 0,
			fmt.Errorf("failed to commit transaction: %w", err)
	}

	return prevBoughtAmount, prevAvgPrice, from_id, nil
}
func (ds *DatabaseService) PositiveBid(update *models.BidUpdate) error {
	prev_amount, prev_avg, from_id, err := ds.UpdateOrderStatus(*update.OrderID, *update.Status, update.BoughtAmount, update.PaidPrice)
	if err != nil {
		return err
	}
	balanceErr := ds.UpdateUserCurrencyBalance(*update.OrderID, strconv.Itoa(from_id),
		*update.PaidPrice-(*prev_amount)*(*prev_avg))
	return balanceErr
}
