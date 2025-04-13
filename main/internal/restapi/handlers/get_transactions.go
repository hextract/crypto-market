package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
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
		dateFrom = (*time.Time)(params.DateFrom)
	}
	if params.DateTo != nil {
		dateTo = (*time.Time)(params.DateTo)
	}

	transfers, err := handler.Database.GetTransfers(user, minAmount, maxAmount, status, currency, operation, dateFrom, dateTo)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	var values []*operations.GetTransactionsTransfersOKBodyItems0
	for _, transfer := range transfers {
		date := strfmt.DateTime(transfer.Date)
		values = append(values, &operations.GetTransactionsTransfersOKBodyItems0{
			ID:         &transfer.ID,
			Currency:   &transfer.Currency,
			Amount:     &transfer.Amount,
			Commission: &transfer.Commission,
			Operation:  &transfer.Operation,
			Status:     &transfer.Status,
			Date:       &date,
			Address:    transfer.Address,
		})
	}

	result := new(operations.GetTransactionsTransfersOK)
	result.SetPayload(values)
	return result
}
