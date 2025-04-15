package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
	"time"
)

func (handler *Handler) GetTransactionsTransfersHandler(params operations.GetTransactionsTransfersParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	var minAmount, maxAmount *float32
	if params.MinAmount != nil {
		minAmount = params.MinAmount
	}
	if params.MaxAmount != nil {
		maxAmount = params.MaxAmount
	}

	var status, currency, operation *string
	if params.Status != nil {
		status = params.Status
	}
	if params.Currency != nil {
		currency = params.Currency
	}
	if params.Operation != nil {
		operation = params.Operation
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

	transfers, err := handler.Database.GetTransfers(user, minAmount, maxAmount, status, currency, operation, dateFrom, dateTo, limit, offset)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	result := new(operations.GetTransactionsTransfersOK)
	result.SetPayload(transfers)
	return result
}
