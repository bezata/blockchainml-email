package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	EmailRequests     *prometheus.CounterVec
	EmailLatency      *prometheus.HistogramVec
	ActiveConnections prometheus.Gauge
	CacheHits         prometheus.Counter
	CacheMisses       prometheus.Counter
	DatabaseLatency   *prometheus.HistogramVec
	R2Latency        *prometheus.HistogramVec
	SearchLatency      *prometheus.HistogramVec
	IndexingLatency    *prometheus.HistogramVec
	CacheLatency       *prometheus.HistogramVec
	NotificationsSent  *prometheus.CounterVec
	WebSocketConns     prometheus.Gauge
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
				Buckets:  prometheus.DefBuckets,
			},
			[]string{"operation"},
		),
		ActiveConnections: promauto.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:     "active_connections",
				Help:     "Number of active connections",
			},
		),
		DatabaseLatency: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:     "database_latency_seconds",
				Help:     "Database operation latency in seconds",
				Buckets:  prometheus.DefBuckets,
			},
			[]string{"operation"},
		),
		SearchLatency: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:     "search_latency_seconds",
				Help:     "Search operation latency in seconds",
			},
			[]string{"operation"},
		),
		IndexingLatency: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:     "indexing_latency_seconds",
				Help:     "Indexing operation latency in seconds",
			},
			[]string{"operation"},
		),
		CacheLatency: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:     "cache_latency_seconds",
				Help:     "Cache operation latency in seconds",
			},
			[]string{"operation"},
		),
		NotificationsSent: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:     "notifications_sent_total",
				Help:     "Total number of notifications sent",
			},
			[]string{"type"},
		),
		WebSocketConns: promauto.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:     "websocket_connections",
				Help:     "Number of WebSocket connections",
			},
		),
	}
}
