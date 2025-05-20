package handlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/main/internal/utils"
	"log"
)

func (h *Handler) UpdateOrderStatusHandler(params operations.UpdateOrderStatusParams) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)
	fmt.Println("Some updates:", len(params.Body))

	for _, bidUpdate := range params.Body {
		bidId, err := h.Database.CheckBid(*bidUpdate.OrderID)
		if bidId == "" || err != nil {
			log.Println("Bid: ", bidId, " Skipped")
			continue
		}
		log.Println("Bid: ", bidId, " continued")

		if *bidUpdate.Status == models.BidStatusCancelled {
			cancelErr := h.Database.CancelBid(bidId)
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
