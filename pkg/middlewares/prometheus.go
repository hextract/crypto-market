package middlewares

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

type PrometheusMetrics struct {
	TotalRequests   *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec
	ErrorResponses  *prometheus.CounterVec
}

func NewPrometheusMetrics() *PrometheusMetrics {
	metrics := &PrometheusMetrics{
		TotalRequests: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Общее количество HTTP-запросов",
			},
			[]string{"method", "endpoint"},
		),
		RequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "Продолжительность обработки запросов (в секундах)",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"method", "endpoint"},
		),
		ErrorResponses: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_error_responses_total",
				Help: "Общее количество ошибок",
			},
			[]string{"endpoint", "status", "error_text"},
		),
	}
	prometheus.MustRegister(metrics.TotalRequests, metrics.RequestDuration, metrics.ErrorResponses)
	return metrics
}

func (metrics *PrometheusMetrics) ApplyMetrics(handler http.Handler) http.Handler {
	return metrics.PrometheusDurationMiddleware(metrics.PrometheusCountMiddleware(metrics.PrometheusErrorMiddleware(handler)))
}
