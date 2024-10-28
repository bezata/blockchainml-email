package cache

import (
	"github.com/bezata/blockchainml-email/internal/monitoring/metrics"
	"go.uber.org/zap"
)

// Cache defines the interface for caching operations
type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Delete(key string) error
}

type cache struct {
	// Define cache fields
}

func NewCache(redis interface{}, cache interface{}, logger *zap.Logger, metrics *metrics.Metrics) *Cache {
	return &cache{}
}
