package handlers

import (
	"github.com/bezata/blockchainml-email/internal/domain/email"
	"github.com/bezata/blockchainml-email/internal/monitoring/metrics"
	"github.com/bezata/blockchainml-email/pkg/realtime"
	"github.com/bezata/blockchainml-email/pkg/search"
	"go.uber.org/zap"
)

type Handlers struct {
    Email    *EmailHandler
    Staff    *StaffHandler
    Thread   *ThreadHandler
    Auth     *AuthHandler
    Realtime *RealtimeHandler
}

func NewHandlers(
    emailService *email.Service,
    searchEngine *search.SearchEngine,
    notifier *realtime.Notifier,
    logger *zap.Logger,
    metrics *metrics.Metrics,
) *Handlers {
    return &Handlers{
        Email:    NewEmailHandler(emailService, searchEngine, notifier, logger, metrics),
        Staff:    NewStaffHandler(staffService, logger, metrics),
        Thread:   NewThreadHandler(threadService, logger, metrics),
        Auth:     NewAuthHandler(authService, logger, metrics),
        Realtime: NewRealtimeHandler(notifier, logger, metrics),
    }
}
