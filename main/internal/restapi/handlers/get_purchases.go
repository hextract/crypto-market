package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
	"time"
)

func (handler *Handler) GetTransactionsPurchaseHandler(params operations.GetTransactionsPurchaseParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	var status *string
	if params.Status != nil {
		status = params.Status
	}

	var dateFrom, dateTo *time.Time
	if params.DateFrom != nil {
		dateFrom = (*time.Time)(params.DateFrom)
	}
	if params.DateTo != nil {
		dateTo = (*time.Time)(params.DateTo)
	}

	purchases, err := handler.Database.GetPurchases(user, status, dateFrom, dateTo)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	var values []*operations.GetTransactionsPurchaseOKBodyItems0
	for _, purchase := range purchases {
		date := strfmt.DateTime(purchase.Date)
		values = append(values, &operations.GetTransactionsPurchaseOKBodyItems0{
			ID:           &purchase.ID,
			CurrencyFrom: &purchase.CurrencyFrom,
			CurrencyTo:   &purchase.CurrencyTo,
			AmountFrom:   &purchase.AmountFrom,
			AmountTo:     &purchase.AmountTo,
			Status:       &purchase.Status,
			Date:         &date,
		})
	}

	result := new(operations.GetTransactionsPurchaseOK)
	result.SetPayload(values)
	return result
}
