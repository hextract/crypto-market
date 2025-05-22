package mappers

import (
	"fetcher/internal/domain/entities"
	"fetcher/internal/presentation/web/dto"
)

func OrderBookSnapshotToApiDto(entity *entities.OrderBookSnapshot) *dto.OrderBookSnapshot {
	dto := &dto.OrderBookSnapshot{}
	dto.AskVolume = &entity.AskVolume
	dto.BidVolume = &entity.BidVolume
	dto.Depth = &entity.Depth
	dto.Mid = &entity.Mid
	dto.Spread = &entity.Spread
	return dto
}
