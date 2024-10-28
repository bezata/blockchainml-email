package services

import (
	"github.com/bezata/blockchainml-mail/internal/config"
	"github.com/bezata/blockchainml-mail/internal/metrics"
	"github.com/bezata/blockchainml-mail/internal/security"
	"github.com/bezata/blockchainml-mail/pkg/cache"
	"github.com/bezata/blockchainml-mail/pkg/realtime"
	"github.com/bezata/blockchainml-mail/pkg/search"
	"go.uber.org/zap"
)

type Services struct {
	Email     *EmailService
	Staff     *StaffService
	Thread    *ThreadService
	Auth      *AuthService
	Search    *search.SearchEngine
	Cache     *cache.Cache
	Notifier  *realtime.Notifier
	Security  *security.Service
}

func New(
	cfg *config.Config,
	cache *cache.Cache,
	searchEngine *search.SearchEngine,
	notifier *realtime.Notifier,
	logger *zap.Logger,
	metrics *metrics.Metrics,
) *Services {
	securityService := security.NewService(cfg.Security, logger)
	
	return &Services{
		Email:    NewEmailService(cfg, cache, searchEngine, logger, metrics),
		Staff:    NewStaffService(cfg, cache, logger, metrics),
		Thread:   NewThreadService(cfg, cache, logger, metrics),
		Auth:     NewAuthService(cfg, securityService, logger, metrics),
		Search:   searchEngine,
		Cache:    cache,
		Notifier: notifier,
		Security: securityService,
	}
}
