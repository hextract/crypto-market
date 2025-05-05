package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
)

func (h *Handler) CancelBidHandler(params operations.CancelBidParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	err := h.Database.CancelBid(params.BidID)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	result := new((operations.CancelBidOK))
	result.SetPayload(&operations.CancelBidOKBody{
		ID:     params.BidID,
		Status: "cancelled",
	})

	err = h.MatchingEngine.CancelOrder(params.BidID)
	if err != nil {
		return utils.HandleError("couldn't cancel order", http.StatusInternalServerError)
	}

	return result
}
