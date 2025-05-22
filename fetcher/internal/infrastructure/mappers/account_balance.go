package mappers

import (
	"fetcher/internal/domain/entities"
	"fetcher/internal/infrastructure/dto"
)

func AccountBalanceToJsonDto(entity *entities.AccountBalance) *dto.JsonAccountBalance {
	dto := &dto.JsonAccountBalance{}
	dto.BtcFree = entity.BtcFree
	dto.BtcLocked = entity.BtcLocked
	dto.EthFree = entity.EthFree
	dto.EthLocked = entity.EthLocked
	dto.UsdtFree = entity.UsdtFree
	dto.UsdtLocked = entity.UsdtLocked
	return dto
}

func AccountBalanceToEntity(dto *dto.JsonAccountBalance) *entities.AccountBalance {
	entity := &entities.AccountBalance{}
	entity.BtcFree = dto.BtcFree
	entity.BtcLocked = dto.BtcLocked
	entity.EthFree = dto.EthFree
	entity.EthLocked = dto.EthLocked
	entity.UsdtFree = dto.UsdtFree
	entity.UsdtLocked = dto.UsdtLocked
	return entity
}
