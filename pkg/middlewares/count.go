package middlewares

import "net/http"

// PrometheusCountMiddleware - middleware для подсчета запросов
func (metrics *PrometheusMetrics) PrometheusCountMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		endpoint := r.URL.Path

		metrics.TotalRequests.WithLabelValues(method, endpoint).Inc()
		next.ServeHTTP(w, r)
	})
}
