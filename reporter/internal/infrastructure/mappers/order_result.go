package mappers

import (
	"reporter/internal/domain/entities"
	"reporter/internal/infrastructure/dto"
)

func OrderResultToEntity(dto *dto.JsonOrderResult) *entities.OrderResult {
	entity := &entities.OrderResult{}
	entity.ExecutionResult = ExecutionResultToEntity(dto.ExecutionResult)
	entity.Order = OrderToEntity(dto.Order)
	return entity
}

func OrderResultToApiDto(entity *entities.OrderResult) *dto.OrderResult {
	dto := &dto.OrderResult{}
	dto.ExecutionResult = ExecutionResultToApiDto(entity.ExecutionResult)
	dto.Order = OrderToApiDto(entity.Order)
	return dto
}
