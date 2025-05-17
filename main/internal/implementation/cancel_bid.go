package implementation

import "context"

func (ds *DatabaseService) CancelBid(ID string) error {
	query := `DELETE FROM bids WHERE bids.id = $1;`

	args := []any{ID}

	rows, err := ds.pool.Query(context.Background(), query, args...)
	rows.Close()
	return err
}
