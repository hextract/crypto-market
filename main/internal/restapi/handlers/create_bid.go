package handlers

import (
	"fmt"
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

	needBalance := *params.Body.MaxPrice * *params.Body.AmountToBuy
	if *params.Body.FromCurrency == "BTC" {
		needBalance = *params.Body.AmountToBuy
	}

	balances, err := h.Database.GetAccountBalance(user)

	quoteBal := float32(0)

	for _, cur := range balances {
		if cur.Currency == *params.Body.FromCurrency {
			quoteBal = cur.Amount
			break
		}
	}

	if quoteBal < needBalance {
		return utils.HandleError("not enough balance!", http.StatusInternalServerError)
	}

	buySpeed := float32(0)
	if params.Body.BuySpeed != nil {
		buySpeed = *params.Body.BuySpeed
	}

	bidID, err := h.Database.CreateBid(
		user.UserID,
		*params.Body.FromCurrency,
		*params.Body.ToCurrency,
		float32(*params.Body.MinPrice),
		float32(*params.Body.MaxPrice),
		float32(*params.Body.AmountToBuy),
		buySpeed,
	)
	if err != nil {
		return utils.HandleInternalError(err)
	}
	updateErr := h.Database.UpdateUserCurrencyBalance(bidID, *params.Body.FromCurrency, -needBalance)
	if updateErr != nil {
		return utils.HandleInternalError(err)
	}

	placeErr := h.MatchingEngine.PlaceOrder(models.Bid{
		ID:           &bidID,
		AmountToBuy:  params.Body.AmountToBuy,
		MinPrice:     params.Body.MinPrice,
		MaxPrice:     params.Body.MaxPrice,
		BuySpeed:     &buySpeed,
		FromCurrency: params.Body.FromCurrency,
		ToCurrency:   params.Body.ToCurrency,
	})
	if placeErr != nil {
		fmt.Println(placeErr)
		cancelErr := h.Database.CancelBid(bidID)
		if cancelErr != nil {
		}
		return utils.HandleError(fmt.Sprintf("couldn't place order, %s", cancelErr), http.StatusInternalServerError)
	}

	result := new(operations.CreateBidOK)
	result.SetPayload(&operations.CreateBidOKBody{
		ID: bidID,
	})
	return result
}
