package mappers

import (
	"fetcher/internal/domain/entities"
	"fetcher/internal/presentation/web/dto"
)

func AccountBalanceToApiDto(entity *entities.AccountBalance) *dto.MyBalance {
	dto := &dto.MyBalance{}
	dto.BtcFree = &entity.BtcFree
	dto.BtcLocked = &entity.BtcLocked
	dto.EthFree = &entity.EthFree
	dto.EthLocked = &entity.EthLocked
	dto.UsdtFree = &entity.UsdtFree
	dto.UsdtLocked = &entity.UsdtLocked
	return dto
}
