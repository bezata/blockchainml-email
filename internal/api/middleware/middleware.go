package middleware

import (
	"github.com/bezata/blockchainml-email/internal/monitoring/metrics"
	"go.uber.org/zap"
)

type Middleware struct {
	Auth      *AuthMiddleware
	RateLimit *RateLimitMiddleware
	Logger    *LoggerMiddleware
	Metrics   *MetricsMiddleware
}

func NewMiddleware(logger *zap.Logger, metrics *metrics.Metrics) *Middleware {
	return &Middleware{
		Auth:      NewAuthMiddleware(logger),
		RateLimit: NewRateLimitMiddleware(logger, metrics),
		Logger:    NewLoggerMiddleware(logger),
		Metrics:   NewMetricsMiddleware(metrics),
	}
}
