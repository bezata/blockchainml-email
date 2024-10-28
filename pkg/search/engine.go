package search

import (
    "context"
    "encoding/json"
    "strings"
    "time"
    "github.com/elastic/go-elasticsearch/v8"
    "github.com/redis/go-redis/v9"
    "go.uber.org/zap"
    "github.com/bezata/blockchainml-email/internal/monitoring/metrics"
)
