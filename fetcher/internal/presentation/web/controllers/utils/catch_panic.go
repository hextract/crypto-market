package utils

import (
	"fetcher/internal/presentation/web/dto"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
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
