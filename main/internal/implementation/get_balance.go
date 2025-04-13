package implementation

import (
	"context"
	"github.com/h4x4d/crypto-market/main/internal/models"
)

func (ds *DatabaseService) GetAccountBalance(user *models.User) ([]*models.UserCurrency, error) {
	query := "select currencies.name, balance from user_balance as b join currencies on currencies.currency_id = b.currency_id where b.user_id = $1;"

	row, errGet := ds.pool.Query(context.Background(), query, user.UserID)
	if errGet != nil {
		return nil, errGet
	}
	defer row.Close()

	result := []*models.UserCurrency{}
	for row.Next() {
		curr := new(models.UserCurrency)

		err := row.Scan(&curr.Currency, &curr.Amount)
		if err != nil {
			return nil, err
		}
		result = append(result, curr)
	}

	return result, nil
}
