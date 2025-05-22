package utils

import (
	"receiver/internal/presentation/web/dto"
)

func GetModelError(emergedError error, code int64) *dto.Error {
	return &dto.Error{
		ErrorMessage:    emergedError.Error(),
		ErrorStatusCode: &code,
	}
}
