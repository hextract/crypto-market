package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
)

func (h *Handler) CreateBidHandler(params operations.CreateBidParams, user *models.User) middleware.Responder {
	if params.Body.FromCurrency == params.Body.ToCurrency {
		return utils.HandleError("from_currency and to_currency must be different", http.StatusInternalServerError)
	}

	if *params.Body.MaxPrice < *params.Body.MinPrice {
		return utils.HandleError("max_price must be greater than or equal to min_price", http.StatusInternalServerError)
	}

	buySpeed := float32(0)
	if params.Body.BuySpeed != nil {
		buySpeed = *params.Body.BuySpeed
	}

	bidID, err := h.Database.CreateBid(
		user.UserID,
		string(*params.Body.FromCurrency),
		string(*params.Body.ToCurrency),
		float32(*params.Body.MinPrice),
		float32(*params.Body.MaxPrice),
		float32(*params.Body.AmountToBuy),
		buySpeed,
	)
	if err != nil {
		return utils.HandleInternalError(err)
	}

	result := new(operations.CreateBidOK)
	result.SetPayload(&operations.CreateBidOKBody{
		ID: bidID,
	})

	err = h.MatchingEngine.PlaceOrder(models.Bid{
		ID: &bidID,
	}) 

	if err != nil {
		return utils.HandleError("couldn't place order", http.StatusInternalServerError)
	}
	
	return result
}
