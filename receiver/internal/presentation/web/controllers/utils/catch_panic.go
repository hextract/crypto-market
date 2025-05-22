package utils

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"receiver/internal/presentation/web/dto"
)

func CatchPanic(responder *middleware.Responder, code int64) {
	if err := recover(); err != nil {
		message := fmt.Sprintf("%v", err)
		*responder = middleware.Error(int(code), dto.Error{
			ErrorMessage:    message,
			ErrorStatusCode: &code,
		})
	}
}
