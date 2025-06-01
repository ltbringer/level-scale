package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	HttpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Estimate RPS.",
	}, []string{"path", "method", "status"})
)

func Init() {
	prometheus.MustRegister(HttpRequestsTotal)
}
