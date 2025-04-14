package handlers

import (
	"github.com/go-openapi/runtime/middleware"
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
		df := time.Unix(*params.DateFrom, 0)
		dateFrom = &df
	}
	if params.DateTo != nil {
		df := time.Unix(*params.DateTo, 0)
		dateTo = &df
	}

	var limit, offset *int64
	if params.Limit != nil {
		limit = params.Limit
	}
	if params.Offset != nil {
		offset = params.Offset
	}

	purchases, err := handler.Database.GetPurchases(user, status, dateFrom, dateTo, limit, offset)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	result := new(operations.GetTransactionsPurchaseOK)
	result.SetPayload(purchases)
	return result
}
