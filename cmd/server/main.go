package main

import (
	"fmt"
	"net/http"

	"github.com/bezata/blockchainml-email/internal/api/handlers"
	"github.com/bezata/blockchainml-email/internal/api/middleware"
	"github.com/bezata/blockchainml-email/internal/api/router"
	"github.com/bezata/blockchainml-email/internal/config"
	"github.com/bezata/blockchainml-email/internal/monitoring/logging"
	"github.com/bezata/blockchainml-email/internal/monitoring/metrics"
	"github.com/bezata/blockchainml-email/internal/services"
	"github.com/bezata/blockchainml-email/pkg/cache"
	"github.com/bezata/blockchainml-email/pkg/realtime"
	"github.com/bezata/blockchainml-email/pkg/search"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger, err := logging.NewLogger()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	defer logger.Sync()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}

	// Initialize metrics
	metrics := metrics.NewMetrics("email_server")

	// Initialize services
	services, err := initializeServices(cfg, logger, metrics)
	if err != nil {
		logger.Fatal("Failed to initialize services", zap.Error(err))
	}

	// Initialize handlers
	handlers := handlers.NewHandlers(services, logger, metrics)

	// Initialize middleware
	middleware := middleware.NewMiddleware(logger, metrics)

	// Initialize router
	r := router.NewRouter(handlers, middleware)

	// Start server
	logger.Info("Starting server", zap.String("port", cfg.Server.Port))
	if err := http.ListenAndServe(":"+cfg.Server.Port, r); err != nil {
		logger.Fatal("Server failed", zap.Error(err))
	}
}

func initializeServices(cfg *config.Config, logger *zap.Logger, metrics *metrics.Metrics) (*services.Services, error) {
	// Initialize cache
	cache := cache.NewCache(cfg.Redis, cfg.Cache, logger, metrics)

	// Initialize search engine
	searchEngine, err := search.NewSearchEngine(cfg.Search, logger, metrics)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize search engine: %w", err)
	}

	// Initialize notifier
	notifier := realtime.NewNotifier(cfg.Redis, logger, metrics)

	// Initialize services
	return services.New(
		cfg,
		cache,
		searchEngine,
		notifier,
		logger,
		metrics,
	), nil
}
