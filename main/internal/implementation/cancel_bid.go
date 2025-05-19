package implementation

import (
	"context"
	"strconv"
)

func (ds *DatabaseService) CancelBid(ID string) error {
	query := `
			UPDATE bids 
			SET 
				status = 'cancelled',
				complete_date = CURRENT_TIMESTAMP
			WHERE id = $1
			AND status IN ('pending', 'processing')
			RETURNING from_id, max_price * amount_to_buy;`

	currencyToReturn := new(int)
	amountToReturn := new(float32)

	err := ds.pool.QueryRow(context.Background(), query, ID).Scan(&currencyToReturn, &amountToReturn)
	if err != nil {
		return err
	}

	updateErr := ds.UpdateUserCurrencyBalance(ID, strconv.Itoa(*currencyToReturn), *amountToReturn)
	return updateErr
}
