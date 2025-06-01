package middleware

import (
	"level-scale/metrics"
	"net/http"
	"strconv"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// MetricsMiddleware Records metrics for each incoming HTTP request.
// Specifically, it increments a counter by labeling it with URL path, request method and HTTP status code.
//
// ResponseWriter is wrapped by responseWriter to capture the status code.
// If we pass ResponseWriter (w) and Request (r) to next.ServeHTTP(w, r)
// ResponseWriter doesn't expose `.status`. So it wouldn't be possible to know
// in case an API returns an error code.
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		start := time.Now()
		next.ServeHTTP(rw, r)
		duration := time.Since(start)
		metrics.HttpRequestsTotal.WithLabelValues(r.URL.Path, r.Method, strconv.Itoa(rw.status)).Inc()
		_ = duration
	})
}
