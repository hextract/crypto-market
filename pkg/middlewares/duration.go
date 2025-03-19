package middlewares

import (
	"net/http"
	"time"
)

// PrometheusDurationMiddleware - измеряет продолжительность обработки запросов
func (metrics *PrometheusMetrics) PrometheusDurationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		endpoint := r.URL.Path
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start).Seconds()
		metrics.RequestDuration.WithLabelValues(method, endpoint).Observe(duration)
	})
}
