package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
)

func (h *Handler) GetBidByID(params operations.GetBidByIDParams, user *models.User) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	bid, err := h.Database.GetBidByID(params.BidID)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	result := new(operations.GetBidByIDOK)
	result.SetPayload(bid)
	return result
}