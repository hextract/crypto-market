package mappers

import (
	"reporter/internal/domain/entities"
	"reporter/internal/infrastructure/dto"
)

func ExecutionResultToEntity(dto *dto.JsonExecutionResult) *entities.ExecutionResult {
	entity := &entities.ExecutionResult{}
	entity.Price = dto.Price
	entity.Volume = dto.Volume
	return entity
}

func ExecutionResultToApiDto(entity *entities.ExecutionResult) *dto.ExecutionResult {
	dto := &dto.ExecutionResult{}
	dto.Price = &entity.Price
	dto.Volume = &entity.Volume
	return dto
}
