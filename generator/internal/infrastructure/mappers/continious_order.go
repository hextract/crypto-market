package mappers

import (
	"generator/internal/domain/entities"
	"generator/internal/infrastructure/dto"
)

func ContinuousOrderToJsonDto(entity *entities.ContinuousOrder) (*dto.ContinuousOrder, error) {
	dto := &dto.ContinuousOrder{}
	dto.Pair = []string{entity.Pair.Base, entity.Pair.Quote}
	dto.OrderID = &entity.OrderID
	dto.Amount = &entity.Amount
	dto.BuySellIndicator = &entity.IsBid
	dto.Speed = &entity.MaxSpeed
	dto.PriceHigh = &entity.PriceHigh
	dto.PriceLow = &entity.PriceLow
	return dto, nil
}

//func ContinuousOrderToEntity(dto *dto.ContinuousOrder) (*entities.ContinuousOrder, error) {
//	entity := &entities.ContinuousOrder{}
//	entity.OrderID = dto.OrderID
//	entity.BuySellIndicator = dto.BuySellIndicator
//	entity.Amount = dto.Amount
//	entity.Pair = dto.Pair
//	entity.Speed = dto.Speed
//	entity.PriceHigh = dto.PriceHigh
//	entity.PriceLow = dto.PriceLow
//	return entity, nil
//}
