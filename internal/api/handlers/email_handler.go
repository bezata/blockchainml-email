package handlers

import (
	"github.com/bezata/blockchainml-email/internal/domain/email"
	"github.com/bezata/blockchainml-email/internal/monitoring/metrics"
	"github.com/bezata/blockchainml-email/pkg/realtime"
	"github.com/bezata/blockchainml-email/pkg/search"
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
