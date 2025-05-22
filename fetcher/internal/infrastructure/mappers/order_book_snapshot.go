package mappers

import (
	"fetcher/internal/domain/entities"
	"fetcher/internal/infrastructure/dto"
)

func OrderBookSnapshotToJsonDto(entity *entities.OrderBookSnapshot) *dto.JsonOrderBookSnapshot {
	dto := &dto.JsonOrderBookSnapshot{}
	dto.AskVolume = entity.AskVolume
	dto.BidVolume = entity.BidVolume
	dto.Depth = entity.Depth
	dto.Mid = entity.Mid
	dto.Spread = entity.Spread
	dto.BestAsk = entity.BestAsk
	dto.BestBid = entity.BestBid
	dto.Imbalance = entity.Imbalance
	dto.Timestamp = entity.Timestamp
	dto.LiqAsk1Pct = entity.LiqAsk1Pct
	dto.LiqBid1Pct = entity.LiqBid1Pct
	dto.VWAPAsk = entity.VWAPAsk
	dto.VWAPBid = entity.VWAPBid
	return dto
}

func OrderBookSnapshotToEntity(dto *dto.JsonOrderBookSnapshot) *entities.OrderBookSnapshot {
	entity := &entities.OrderBookSnapshot{}
	entity.AskVolume = dto.AskVolume
	entity.BidVolume = dto.BidVolume
	entity.Depth = dto.Depth
	entity.Mid = dto.Mid
	entity.Spread = dto.Spread
	entity.BestAsk = dto.BestAsk
	entity.BestBid = dto.BestBid
	entity.Imbalance = dto.Imbalance
	entity.Timestamp = dto.Timestamp
	entity.LiqAsk1Pct = dto.LiqAsk1Pct
	entity.LiqBid1Pct = dto.LiqBid1Pct
	entity.VWAPAsk = dto.VWAPAsk
	entity.VWAPBid = dto.VWAPBid
	return entity
}
