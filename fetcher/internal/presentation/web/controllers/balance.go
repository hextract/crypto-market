package controllers

import (
	"fetcher/internal/application/use_cases"
	"fetcher/internal/presentation/web/controllers/utils"
	"fetcher/internal/presentation/web/mappers"
	"fetcher/internal/presentation/web/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

type AccountBalanceController struct {
	getAccountBalanceUseCase *use_cases.GetAccountBalanceUseCase
}

func (controller *AccountBalanceController) GetAccountBalance(params operations.GetBalanceParams) (responder middleware.Responder) {
	internalErrorCode := int64(operations.GetBalanceInternalServerErrorCode)
	defer utils.CatchPanic(&responder, internalErrorCode)

	balance, errBalance := controller.getAccountBalanceUseCase.Execute()
	if errBalance != nil {
		result := &operations.GetBalanceInternalServerError{}
		result.SetPayload(utils.GetModelError(errBalance, internalErrorCode))
		return result
	}

	apiAccountBalance := mappers.AccountBalanceToApiDto(balance)
	result := &operations.GetBalanceOK{}
	result.SetPayload(apiAccountBalance)
	return result
}

func NewAccountBalanceController(getAccountBalanceUseCase *use_cases.GetAccountBalanceUseCase) (*AccountBalanceController, error) {
	accountBalanceController := &AccountBalanceController{}
	accountBalanceController.getAccountBalanceUseCase = getAccountBalanceUseCase
	return accountBalanceController, nil
}
