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

func (h *Handler) LoginHandler(api operations.PostAuthLoginParams) middleware.Responder {
	// Tracing
	_, span := h.tracer.Start(context.Background(), "login")
	defer span.End()

	token, err := impl.LoginUser(h.Client, api.Body)
	var apiErr *gocloak.APIError
	if errors.As(err, &apiErr) {
		// Logging
		slog.Error(
			"failed login user",
			slog.String("method", "POST"),
			slog.Group("user-properties",
				slog.String("login", api.Body.Login),
			),
			slog.Int("status_code", operations.PostAuthLoginUnauthorizedCode),
			slog.String("error", err.Error()),
		)

		switch apiErr.Code {
		case 401:
			return new(operations.PostAuthLoginUnauthorized).WithPayload(&models.Error{
				ErrorMessage:    err.Error(),
				ErrorStatusCode: int64(operations.PostAuthLoginUnauthorizedCode),
			})
		default:
			return new(operations.PostAuthLoginInternalServerError).WithPayload(&models.Error{
				ErrorMessage:    err.Error(),
				ErrorStatusCode: int64(operations.PostAuthLoginInternalServerErrorCode),
			})
		}
	} else if err != nil {
		return new(operations.PostAuthLoginInternalServerError).WithPayload(&models.Error{
			ErrorMessage:    err.Error(),
			ErrorStatusCode: int64(operations.PostAuthLoginInternalServerErrorCode),
		})
	}

	// Logging
	slog.Info(
		"user login",
		slog.String("method", "POST"),
		slog.Group("user-properties",
			slog.String("login", api.Body.Login),
		),
		slog.Int("status_code", operations.PostAuthLoginOKCode),
	)

	result := new(operations.PostAuthLoginOK).WithPayload(&operations.PostAuthLoginOKBody{
		Token: *token,
	})
	return result
}
