package implementation

import "context"

func (ds *DatabaseService) CancelBid(ID string) error {
	query := `
			UPDATE bids 
			SET 
				status = 'cancelled',
				complete_date = CURRENT_TIMESTAMP
			WHERE id = $1
			AND status IN ('pending', 'processing');`

	args := []any{ID}

	rows, err := ds.pool.Query(context.Background(), query, args...)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}
