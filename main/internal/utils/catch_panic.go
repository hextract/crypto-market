package utils

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func CatchPanic(responder *middleware.Responder) {
	if err := recover(); err != nil {
		errText := fmt.Sprintf("%v", err)
		*responder = HandleError(errText, http.StatusInternalServerError)
	}
}
