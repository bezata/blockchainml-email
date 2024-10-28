package handlers

import (
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/bezata/blockchainml-mail/internal/domain/email"
    "github.com/bezata/blockchainml-mail/pkg/search"
    "github.com/bezata/blockchainml-mail/pkg/realtime"
    "github.com/bezata/blockchainml-mail/internal/monitoring/metrics"
    "go.uber.org/zap"
)

type EmailHandler struct {
    emailService *email.Service
    searchEngine *search.SearchEngine
    notifier    *realtime.Notifier
    logger      *zap.Logger
    metrics     *metrics.Metrics
}

func NewEmailHandler(
    emailService *email.Service,
    searchEngine *search.SearchEngine,
    notifier *realtime.Notifier,
    logger *zap.Logger,
    metrics *metrics.Metrics,
) *EmailHandler {
    return &EmailHandler{
        emailService: emailService,
        searchEngine: searchEngine,
        notifier:    notifier,
        logger:      logger,
        metrics:     metrics,
    }
}
