package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
)

func (h *Handler) UpdateOrderStatusHandler(params operations.UpdateOrderStatusParams) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	for _, bidUpdate := range params.Body {
		if *bidUpdate.Status == models.BidUpdateStatusCancelled {
			cancelErr := h.Database.CancelBid(*bidUpdate.OrderID)
			if cancelErr != nil {
				return utils.HandleInternalError(cancelErr)
			}
		} else {
			positiveErr := h.Database.PositiveBid(bidUpdate)
			if positiveErr != nil {
				return utils.HandleInternalError(positiveErr)
			}
		}
	}
	result := new(operations.UpdateOrderStatusOK)
	payload := operations.UpdateOrderStatusOKBody{Status: "ok"}
	result.SetPayload(&payload)
	return result
}
