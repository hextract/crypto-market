package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
)

func (handler *Handler) GetBalanceHandler(params operations.GetAccountBalanceParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	value := float32(0.0)

	var _ []*operations.GetAccountBalanceOKBodyItems0

	v := operations.GetAccountBalanceOKBodyItems0{
		Amount:   &value,
		Currency: operations.GetAccountBalanceOKBodyItems0CurrencyBTC,
	}

	result := new(operations.GetAccountBalanceOK)
	values := []*operations.GetAccountBalanceOKBodyItems0{
		&v,
		&v,
	}

	result.SetPayload(values)
	return
}
