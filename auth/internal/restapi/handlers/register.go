package handlers

import (
	"context"
	"errors"
	"github.com/Nerzal/gocloak/v13"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/auth/internal/impl"
	"github.com/h4x4d/crypto-market/auth/internal/models"
	"github.com/h4x4d/crypto-market/auth/internal/restapi/operations"
	"log/slog"
)

func (h *Handler) RegisterHandler(api operations.PostAuthRegisterParams) middleware.Responder {
	// Tracing
	_, span := h.tracer.Start(context.Background(), "register")
	defer span.End()

	token, err := impl.CreateUser(h.Client, api.Body)

	var apiErr *gocloak.APIError
	if errors.As(err, &apiErr) {
		// Logging
		slog.Error(
			"failed register new user",
			slog.String("method", "POST"),
			slog.Group("user-properties",
				slog.String("login", api.Body.Login),
				slog.String("email", api.Body.Email),
			),
			slog.Int("status_code", apiErr.Code),
			slog.String("error", apiErr.Error()),
		)

		switch apiErr.Code {
		case 409:
			conflict := int64(operations.PostAuthRegisterConflictCode)
			return new(operations.PostAuthRegisterConflict).WithPayload(&models.Error{
				ErrorMessage:    err.Error(),
				ErrorStatusCode: conflict,
			})
		default:
			return new(operations.PostAuthRegisterInternalServerError).WithPayload(&models.Error{
				ErrorMessage:    err.Error(),
				ErrorStatusCode: int64(operations.PostAuthRegisterInternalServerErrorCode),
			})
		}
	} else if err != nil {
		return new(operations.PostAuthRegisterInternalServerError).WithPayload(&models.Error{
			ErrorMessage:    err.Error(),
			ErrorStatusCode: int64(operations.PostAuthRegisterInternalServerErrorCode),
		})
	}
	// Logging
	slog.Info(
		"register new user",
		slog.String("method", "POST"),
		slog.Group("user-properties",
			slog.String("login", api.Body.Login),
			slog.String("email", api.Body.Email),
		),
		slog.Int("status_code", operations.PostAuthRegisterOKCode),
	)

	result := new(operations.PostAuthRegisterOK).WithPayload(&operations.PostAuthRegisterOKBody{
		Token: *token,
	})
	return result
}
