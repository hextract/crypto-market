package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/auth/internal/impl"
	"github.com/h4x4d/crypto-market/auth/internal/models"
	"github.com/h4x4d/crypto-market/auth/internal/restapi/operations"
	"log/slog"
)

func (h *Handler) ValidateTokenHandler(api operations.PostAuthValidateTokenParams) middleware.Responder {
	// Tracing
	_, span := h.tracer.Start(context.Background(), "change_password")
	defer span.End()

	isActive, err := impl.ValidateToken(h.Client, api.Body)
	conflict := int64(operations.PostAuthValidateTokenBadRequestCode)
	if err != nil {
		// Logging
		slog.Error(
			"failed to validate token",
			slog.String("method", "POST"),
			slog.Group("user-properties",
				slog.String("login", api.Body.Token),
			),
			slog.Int("status_code", operations.PostAuthValidateTokenBadRequestCode),
			slog.String("error", err.Error()),
		)

		return new(operations.PostAuthValidateTokenBadRequest).WithPayload(&models.Error{
			ErrorMessage:    err.Error(),
			ErrorStatusCode: conflict,
		})
	}
	// Logging
	slog.Info(
		"user validated token",
		slog.String("method", "POST"),
		slog.Int("status_code", operations.PostAuthValidateTokenOKCode),
	)

	result := new(operations.PostAuthValidateTokenOK).WithPayload(&operations.PostAuthValidateTokenOKBody{
		IsValid: *isActive,
	})
	return result
}
