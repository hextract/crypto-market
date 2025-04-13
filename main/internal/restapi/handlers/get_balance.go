package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
)

func (handler *Handler) GetBalanceHandler(_ operations.GetAccountBalanceParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	balances, err := handler.Database.GetAccountBalance(user)

	if err != nil {
		return utils.HandleInternalError(err)
	}

	var values []*operations.GetAccountBalanceOKBodyItems0

	for _, balance := range balances {
		values = append(values, &operations.GetAccountBalanceOKBodyItems0{
			Currency: balance.Currency,
			Amount:   &balance.Amount,
		})
	}
	result := new(operations.GetAccountBalanceOK)
	result.SetPayload(values)
	return result
}
