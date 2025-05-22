package mappers

import (
	"fetcher/internal/domain/entities"
	"fetcher/internal/presentation/web/dto"
)

func OrderBookDataToApiDto(entity *entities.OrderBookData) *dto.OrderBookData {
	dto := &dto.OrderBookData{}
	for _, entitySnapshot := range entity.Snapshots {
		dtoSnapshot := OrderBookSnapshotToApiDto(entitySnapshot)
		dto.Snapshots = append(dto.Snapshots, dtoSnapshot)
	}
	return dto
}
