package utils

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/main/internal/models"
	"net/http"
)

func HandleInternalError(err error) middleware.Responder {
	errorText := fmt.Sprintf("Internal error occured: %s", err)
	return HandleError(errorText, http.StatusInternalServerError)
}

func HandleError(message string, code int) middleware.Responder {
	properCode := int64(code)
	return middleware.Error(code, &models.Error{
		ErrorMessage:    &message,
		ErrorStatusCode: &properCode,
	})
}
