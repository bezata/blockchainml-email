package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/bezata/blockchainml-email/internal/api/handlers"
    "github.com/bezata/blockchainml-email/internal/api/middleware"
    "github.com/bezata/blockchainml-email/internal/api/router"
    "github.com/bezata/blockchainml-email/internal/config"
    "github.com/bezata/blockchainml-email/internal/monitoring/metrics"
    "github.com/bezata/blockchainml-email/internal/services"
    "github.com/bezata/blockchainml-email/internal/storage"
    "github.com/bezata/blockchainml-email/pkg/cache"
    "github.com/bezata/blockchainml-email/pkg/realtime"
    "github.com/bezata/blockchainml-email/pkg/search"
    "go.mongodb.org/mongo-driver/mongo"
    "go.uber.org/zap"
)

func main() {
    // Initialize context
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Initialize logger
    logger, err := zap.NewProduction()
    if err != nil {
        panic(err)
    }
    defer logger.Sync()

    // Load configuration
    cfg, err := config.Load()
    if err != nil {
        logger.Fatal("Failed to load configuration", zap.Error(err))
    }

    // Initialize metrics
    metrics := metrics.NewMetrics("email_server")

    // Initialize core dependencies
    deps, err := initializeDependencies(ctx, cfg, logger, metrics)
    if err != nil {
        logger.Fatal("Failed to initialize dependencies", zap.Error(err))
    }
    defer deps.cleanup()

    // Initialize services
    services := initializeServices(cfg, deps, logger, metrics)

    // Initialize API components
    apiHandlers := handlers.NewHandlers(services, logger, metrics)  // Pass the entire services struct
    mw := middleware.NewMiddleware(logger, metrics)
    r := router.NewRouter(apiHandlers, mw)

    // Create server
    srv := &http.Server{
        Addr:         ":" + cfg.Server.Port,
        Handler:      r,
        ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
        WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
    }

    // Start server with graceful shutdown
    go func() {
        logger.Info("Starting server", zap.String("port", cfg.Server.Port))
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Fatal("Server failed", zap.Error(err))
        }
    }()

    // Wait for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    // Gracefully shutdown
    logger.Info("Shutting down server...")
    shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 10*time.Second)
    defer shutdownCancel()

    if err := srv.Shutdown(shutdownCtx); err != nil {
        logger.Error("Server forced to shutdown", zap.Error(err))
    }

    logger.Info("Server exited properly")
}

type dependencies struct {
    db          *mongo.Database
    cache       *cache.Cache
    search      *search.SearchEngine
    notifier    *realtime.Notifier
    cleanup     func()
}

func initializeDependencies(ctx context.Context, cfg *config.Config, logger *zap.Logger, metrics *metrics.Metrics) (*dependencies, error) {
    // Initialize MongoDB
    db, err := storage.ConnectMongoDB(ctx, cfg.MongoDB)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
    }

    // Initialize Redis-based cache
    cache, err := cache.NewCache(cfg.Redis, cfg.Cache, logger, metrics)
    if err != nil {
        return nil, fmt.Errorf("failed to initialize cache: %w", err)
    }

    // Initialize search engine
    searchEngine, err := search.NewSearchEngine(cfg.Search, logger, metrics)
    if err != nil {
        return nil, fmt.Errorf("failed to initialize search engine: %w", err)
    }

    // Initialize real-time notifier
    notifier := realtime.NewNotifier(cfg.Redis, logger, metrics)

    // Create cleanup function
    cleanup := func() {
        if err := db.Client().Disconnect(ctx); err != nil {
            logger.Error("Failed to disconnect MongoDB", zap.Error(err))
        }
        if err := cache.Close(); err != nil {
            logger.Error("Failed to close cache", zap.Error(err))
        }
        if err := searchEngine.Close(); err != nil {
            logger.Error("Failed to close search engine", zap.Error(err))
        }
        notifier.Close()
    }

    return &dependencies{
        db:       db,
        cache:    cache,
        search:   searchEngine,
        notifier: notifier,
        cleanup:  cleanup,
    }, nil
}

func initializeServices(
    cfg *config.Config,
    deps *dependencies,
    logger *zap.Logger,
    metrics *metrics.Metrics,
) *services.Services {
    // Initialize storage repositories
    repositories := storage.NewRepositories(deps.db)

    // Initialize services
    return services.New(services.Config{
        Repositories: repositories,
        Cache:       deps.cache,
        Search:      deps.search,
        Notifier:    deps.notifier,
        Config:      cfg,
        Logger:      logger,
        Metrics:     metrics,
    })
}
