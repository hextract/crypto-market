package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"receiver/internal/application/use_cases"
	"receiver/internal/domain/ports/repositories"
	"receiver/internal/presentation/web/controllers/utils"
	"receiver/internal/presentation/web/mappers"
	"receiver/internal/presentation/web/restapi/operations"
)

type OrdersController struct {
	receiveOrderUseCase *use_cases.ReceiveOrderUseCase
}

func (controller *OrdersController) ReceiveOrder(params operations.ReceiveOrderParams) (responder middleware.Responder) {
	internalErrorCode := int64(operations.ReceiveOrderInternalServerErrorCode)
	defer utils.CatchPanic(&responder, internalErrorCode)

	errReceiveOrder := controller.receiveOrderUseCase.Execute(mappers.OrderToEntity(params.Object))
	if errReceiveOrder != nil {
		result := &operations.ReceiveOrderInternalServerError{}
		result.SetPayload(utils.GetModelError(errReceiveOrder, internalErrorCode))
		return result
	}

	result := &operations.ReceiveOrderOK{}
	return result
}

func NewOrdersController(repository repositories.IOrdersRepository) (*OrdersController, error) {
	receiveOrderUseCase, errUseCase := use_cases.NewReceiveOrderUseCase(repository)
	if errUseCase != nil {
		return nil, errUseCase
	}
	return &OrdersController{receiveOrderUseCase: receiveOrderUseCase}, nil
}
