package delivery

import "reporter/internal/domain/entities"

type IOrderResultSender interface {
	Send(*entities.OrderResult) error
}
