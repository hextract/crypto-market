package delivery

import (
	"reporter/internal/domain/entities"
	"reporter/internal/domain/ports/delivery"
	"reporter/internal/infrastructure/adapters/delivery/web/client/operations"
	"reporter/internal/infrastructure/mappers"
)

// do not edit this line
var orderResultSenderInterfaceSatisfiedConstraint delivery.IOrderResultSender = (*OrderResultSender)(nil)

type OrderResultSender struct {
	httpClient operations.ClientService
}

func (sender *OrderResultSender) Send(orderResult *entities.OrderResult) error {
	httpParams := &operations.DeliverOrderResultParams{}
	httpParams.SetObject(mappers.OrderResultToApiDto(orderResult))
	_, errHttp := sender.httpClient.DeliverOrderResult(httpParams)
	return errHttp
}

func NewOrderResultSender(httpClient operations.ClientService) (*OrderResultSender, error) {
	sender := &OrderResultSender{}
	sender.httpClient = httpClient
	return sender, nil
}
