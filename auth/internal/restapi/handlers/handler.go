package handlers

import (
	"github.com/h4x4d/crypto-market/pkg/client"
	"github.com/h4x4d/crypto-market/pkg/jaeger"
	"go.opentelemetry.io/otel/trace"
	"log"
)

type Handler struct {
	Client *client.Client
	tracer trace.Tracer
}

func NewHandler() (*Handler, error) {
	cl, err := client.NewClient()
	if err != nil {
		return nil, err
	}

	tracer, err := jaeger.InitTracer("Auth")
	if err != nil {
		log.Fatal("init tracer", err)
	}

	return &Handler{cl, tracer}, nil
}
