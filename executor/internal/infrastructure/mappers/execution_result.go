package mappers

import (
	"executor/internal/domain/entities"
	"executor/internal/infrastructure/dto"
)

func ExecutionResultToEntity(dto *dto.JsonExecutionResult) *entities.ExecutionResult {
	entity := &entities.ExecutionResult{}
	entity.Price = dto.Price
	entity.Volume = dto.Volume
	entity.Order = OrderToEntity(dto.Order)
	return entity
}

func ExecutionResultToJsonDto(entity *entities.ExecutionResult) *dto.JsonExecutionResult {
	dto := &dto.JsonExecutionResult{}
	dto.Price = entity.Price
	dto.Volume = entity.Volume
	dto.Order = OrderToJsonDto(entity.Order)
	return dto
}
