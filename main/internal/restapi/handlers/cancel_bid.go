package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
)

func (h *Handler) CancelBidHandler(params operations.CancelBidParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	result := new(operations.CancelBidOK)
	result.SetPayload(&operations.CancelBidOKBody{
		ID:     params.BidID,
		Status: "cancelled",
	})

	// TODO UNCOMMENT WHEN MATCHING ENGINE IS ADDED
	// err := h.MatchingEngine.CancelOrder(params.BidID)
	// if err != nil {
	// 	return utils.HandleError("couldn't cancel order", http.StatusInternalServerError)
	// }

	err := h.Database.CancelBid(params.BidID)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	return result
}
