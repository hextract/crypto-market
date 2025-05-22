package controllers

import (
	"fetcher/internal/application/use_cases"
	"fetcher/internal/presentation/web/controllers/utils"
	"fetcher/internal/presentation/web/mappers"
	"fetcher/internal/presentation/web/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

type MarketDataController struct {
	getOrderBookDataUseCase *use_cases.GetOrderBookDataUseCase
}

func (controller *MarketDataController) GetOrderBookData(params operations.GetLastMarketDataParams) (responder middleware.Responder) {
	internalErrorCode := int64(operations.GetLastMarketDataInternalServerErrorCode)
	defer utils.CatchPanic(&responder, internalErrorCode)

	orderBookData, errOrderBookData := controller.getOrderBookDataUseCase.Execute()
	if errOrderBookData != nil {
		result := &operations.GetLastMarketDataInternalServerError{}
		result.SetPayload(utils.GetModelError(errOrderBookData, internalErrorCode))
		return result
	}

	apiOrderBookData := mappers.OrderBookDataToApiDto(orderBookData)
	result := &operations.GetLastMarketDataOK{}
	result.SetPayload(apiOrderBookData)
	return result
}

func NewMarketDataController(getOrderBookDataUseCase *use_cases.GetOrderBookDataUseCase) (*MarketDataController, error) {
	marketDataController := &MarketDataController{}
	marketDataController.getOrderBookDataUseCase = getOrderBookDataUseCase
	return marketDataController, nil
}
