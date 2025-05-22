package mappers

import (
	"receiver/internal/domain/entities"
	"receiver/internal/infrastructure/dto"
)

func OrderToJsonDto(entity *entities.Order) *dto.JsonOrder {
	dto := &dto.JsonOrder{}
	dto.Indicator = entity.Indicator
	dto.Price = entity.Price
	dto.Ticker = entity.Ticker
	dto.Type = entity.Type
	dto.UserID = entity.UserID
	dto.Volume = entity.Volume
	return dto
}
