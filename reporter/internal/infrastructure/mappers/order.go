package mappers

import (
	"reporter/internal/domain/entities"
	"reporter/internal/infrastructure/dto"
)

func OrderToEntity(dto *dto.JsonOrder) *entities.Order {
	entity := &entities.Order{}
	entity.Indicator = dto.Indicator
	entity.Price = dto.Price
	entity.Ticker = dto.Ticker
	entity.Type = dto.Type
	entity.UserID = dto.UserID
	entity.Volume = dto.Volume
	return entity
}

func OrderToApiDto(entity *entities.Order) *dto.Order {
	dto := &dto.Order{}
	dto.Indicator = &entity.Indicator
	dto.Price = &entity.Price
	dto.Ticker = &entity.Ticker
	dto.Type = &entity.Type
	dto.UserID = &entity.UserID
	dto.Volume = &entity.Volume
	return dto
}
