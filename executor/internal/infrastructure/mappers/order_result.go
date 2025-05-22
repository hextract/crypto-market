package mappers

import (
	"executor/internal/domain/entities"
	"executor/internal/infrastructure/dto"
)

func OrderResultToEntity(dto *dto.JsonOrderResult) *entities.OrderResult {
	entity := &entities.OrderResult{}
	entity.ExecutionResult = ExecutionResultToEntity(dto.ExecutionResult)
	entity.Order = OrderToEntity(dto.Order)
	return entity
}

func OrderResultToJsonDto(entity *entities.OrderResult) *dto.JsonOrderResult {
	dto := &dto.JsonOrderResult{}
	dto.ExecutionResult = ExecutionResultToJsonDto(entity.ExecutionResult)
	dto.Order = OrderToJsonDto(entity.Order)
	return dto
}
