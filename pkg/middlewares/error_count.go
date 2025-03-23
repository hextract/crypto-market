package middlewares

import (
	"net/http"
)

// responseWriterWrapper - для захвата статуса ответа
type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// PrometheusErrorMiddleware - отслеживает ошибки
func (metrics *PrometheusMetrics) PrometheusErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		endpoint := r.URL.Path

		wrapper := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapper, r)

		if wrapper.statusCode >= 400 {
			metrics.ErrorResponses.WithLabelValues(method, endpoint, http.StatusText(wrapper.statusCode)).Inc()
		}
	})
}
