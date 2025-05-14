package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
)

func (handler *Handler) PostTransactionsDepositHandler(params operations.PostTransactionsDepositParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	currency := params.Body.Currency
	transaction, err := handler.Database.Deposit(params.HTTPRequest.Context(), user.UserID, *currency)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	result := new(operations.PostTransactionsDepositOK)
	result.SetPayload(transaction)
	return result
}
