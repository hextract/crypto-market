package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
)

// PostTransactionsWithdrawHandler handles POST /transactions/withdraw requests.
func (handler *Handler) PostTransactionsWithdrawHandler(params operations.PostTransactionsWithdrawParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	// Извлечение параметров из запроса
	currency := params.Body.Currency
	amount := params.Body.Amount
	address := params.Body.Address

	// Вызов сервиса
	transaction, err := handler.Database.Withdraw(params.HTTPRequest.Context(), user.UserID, *currency, *address, *amount)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	// Формирование ответа
	result := new(operations.PostTransactionsWithdrawOK)
	result.SetPayload(transaction)
	return result
}
