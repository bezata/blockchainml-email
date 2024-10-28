package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
    EmailRequests     *prometheus.CounterVec
    EmailLatency      *prometheus.HistogramVec
    DatabaseLatency   *prometheus.HistogramVec
    CacheHits         prometheus.Counter
    CacheMisses       prometheus.Counter
}

func NewMetrics(namespace string) *Metrics {
    return &Metrics{
        EmailRequests: promauto.NewCounterVec(
            prometheus.CounterOpts{
                Namespace: namespace,
                Name:     "email_requests_total",
                Help:     "Total number of email requests",
            },
            []string{"operation", "status"},
        ),
        EmailLatency: promauto.NewHistogramVec(
            prometheus.HistogramOpts{
                Namespace: namespace,
                Name:     "email_latency_seconds",
                Help:     "Email operation latency in seconds",
            },
            []string{"operation"},
        ),
        // Add other metrics...
    }
}
