package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
)

func (h *Handler) UpdateOrderStatusHandler(params operations.UpdateOrderStatusParams, principal interface{}) middleware.Responder {
	if *params.Body.Status == "cancelled" {
		order_id, err := h.Database.UpdateOrderStatus(params.OrderID, *params.Body.Status, nil)
		if err != nil {
			return utils.HandleInternalError(err)
		}
		result := new(operations.CreateBidOK)
		result.SetPayload(&operations.CreateBidOKBody{
			ID: order_id,
		})
		return result
	}

	bid, err := h.Database.GetBidByID(params.OrderID)
	if err != nil {
		return utils.HandleInternalError(err)
	}
	totalPriceTo := (*params.Body.BoughtAmount)
	h.Database.UpdateUserCurrencyBalance(params.OrderID, *bid.ToCurrency, totalPriceTo)

	totalPriceFrom := (*bid.MaxPrice - *params.Body.Price) * (*params.Body.BoughtAmount)
	h.Database.UpdateUserCurrencyBalance(params.OrderID, *bid.FromCurrency, totalPriceFrom)

	order_id, err := h.Database.UpdateOrderStatus(params.OrderID, *params.Body.Status, params.Body.BoughtAmount)
	if err != nil {
		return utils.HandleInternalError(err)
	}
	result := new(operations.CreateBidOK)
	result.SetPayload(&operations.CreateBidOKBody{
		ID: order_id,
	})
	return result
}
