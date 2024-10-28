package middleware

import (
	"github.com/bezata/blockchainml-mail/internal/monitoring/metrics"
	"go.uber.org/zap"
)

type AuthMiddleware struct {
	logger *zap.Logger
}

type RateLimitMiddleware struct {
	logger  *zap.Logger
	metrics *metrics.Metrics
}

type LoggerMiddleware struct {
	logger *zap.Logger
}

type MetricsMiddleware struct {
	metrics *metrics.Metrics
}

func NewAuthMiddleware(logger *zap.Logger) *AuthMiddleware {
	return &AuthMiddleware{logger: logger}
}

func NewRateLimitMiddleware(logger *zap.Logger, metrics *metrics.Metrics) *RateLimitMiddleware {
	return &RateLimitMiddleware{logger: logger, metrics: metrics}
}

func NewLoggerMiddleware(logger *zap.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{logger: logger}
}

func NewMetricsMiddleware(metrics *metrics.Metrics) *MetricsMiddleware {
	return &MetricsMiddleware{metrics: metrics}
}
