package implementation

import (
	"context"
	"fmt"
	"github.com/h4x4d/crypto-market/main/internal/models"
)

// FREE MONEY!!! NOT FOR PRODUCTION

func (ds *DatabaseService) GetAccountBalance(user *models.User) ([]*models.UserCurrency, error) {
	query := "select currencies.name, balance, b.currency_id from user_balance as b join currencies on currencies.currency_id = b.currency_id where b.user_id = $1;"

	row, errGet := ds.pool.Query(context.Background(), query, user.UserID)
	if errGet != nil {
		return nil, errGet
	}
	defer row.Close()

	result := []*models.UserCurrency{}
	currenciesToCheck := map[int]bool{1: true, 2: true}
	for row.Next() {
		curr := new(models.UserCurrency)
		id := 0

		err := row.Scan(&curr.Currency, &curr.Amount, &id)
		if err != nil {
			return nil, err
		}
		currenciesToCheck[id] = false
		result = append(result, curr)
	}

	for id, _ := range currenciesToCheck {
		if currenciesToCheck[id] {
			amount := 100000
			if id == 2 {
				amount = 1
			}

			_, err := ds.pool.Exec(context.Background(),
				"INSERT INTO user_balance (user_id, currency_id, balance) VALUES ($1, $2, $3) "+
					"ON CONFLICT (user_id, currency_id) DO UPDATE SET balance = $3",
				user.UserID, id, amount)

			if err != nil {
				return nil, fmt.Errorf("failed to update balance: %v", err)
			}
		}
	}

	return result, nil
}

//func (ds *DatabaseService) GetAccountBalance(user *models.User) ([]*models.UserCurrency, error) {
//	query := "select currencies.name, balance from user_balance as b join currencies on currencies.currency_id = b.currency_id where b.user_id = $1;"
//
//	row, errGet := ds.pool.Query(context.Background(), query, user.UserID)
//	if errGet != nil {
//		return nil, errGet
//	}
//	defer row.Close()
//
//	result := []*models.UserCurrency{}
//	for row.Next() {
//		curr := new(models.UserCurrency)
//
//		err := row.Scan(&curr.Currency, &curr.Amount)
//		if err != nil {
//			return nil, err
//		}
//		result = append(result, curr)
//	}
//
//	return result, nil
//}
