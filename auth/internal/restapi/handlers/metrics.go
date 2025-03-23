package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/crypto-market/auth/internal/restapi/operations"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsHandler(p operations.GetMetricsParams) middleware.Responder {
	return NewCustomResponder(p.HTTPRequest, promhttp.Handler())
}
