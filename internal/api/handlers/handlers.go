package handlers

import (
    "github.com/bezata/blockchainml-email/internal/services"
    "go.uber.org/zap"
    "github.com/bezata/blockchainml-email/internal/monitoring/metrics"
)

type Handlers struct {
    Email     *EmailHandler
    Staff     *StaffHandler
    Thread    *ThreadHandler
    Auth      *AuthHandler
    Realtime  *RealtimeHandler
}

func NewHandlers(services *services.Services, logger *zap.Logger, metrics *metrics.Metrics) *Handlers {
    return &Handlers{
        Email:    NewEmailHandler(services.Email, logger, metrics),
        Staff:    NewStaffHandler(services.Staff, logger, metrics),
        Thread:   NewThreadHandler(services.Email, logger, metrics),
        Auth:     NewAuthHandler(services.Auth, logger, metrics),
        Realtime: NewRealtimeHandler(services.Email, logger, metrics),
    }
}

// internal/api/handlers/email_handler.go
type EmailHandler struct {
    emailService *services.EmailService
    logger       *zap.Logger
    metrics      *metrics.Metrics
}

func NewEmailHandler(emailService *services.EmailService, logger *zap.Logger, metrics *metrics.Metrics) *EmailHandler {
    return &EmailHandler{
        emailService: emailService,
        logger:      logger,
        metrics:     metrics,
    }
}

// internal/api/handlers/staff_handler.go
type StaffHandler struct {
    staffService *services.StaffService
    logger       *zap.Logger
    metrics      *metrics.Metrics
}

func NewStaffHandler(staffService *services.StaffService, logger *zap.Logger, metrics *metrics.Metrics) *StaffHandler {
    return &StaffHandler{
        staffService: staffService,
        logger:      logger,
        metrics:     metrics,
    }
}

// internal/api/handlers/thread_handler.go
type ThreadHandler struct {
    emailService *services.EmailService
    logger       *zap.Logger
    metrics      *metrics.Metrics
}

func NewThreadHandler(emailService *services.EmailService, logger *zap.Logger, metrics *metrics.Metrics) *ThreadHandler {
    return &ThreadHandler{
        emailService: emailService,
        logger:      logger,
        metrics:     metrics,
    }
}

// internal/api/handlers/auth_handler.go
type AuthHandler struct {
    authService *services.AuthService
    logger      *zap.Logger
    metrics     *metrics.Metrics
}

func NewAuthHandler(authService *services.AuthService, logger *zap.Logger, metrics *metrics.Metrics) *AuthHandler {
    return &AuthHandler{
        authService: authService,
        logger:     logger,
        metrics:    metrics,
    }
}

// internal/api/handlers/realtime_handler.go
type RealtimeHandler struct {
    emailService *services.EmailService
    logger       *zap.Logger
    metrics      *metrics.Metrics
}

func NewRealtimeHandler(emailService *services.EmailService, logger *zap.Logger, metrics *metrics.Metrics) *RealtimeHandler {
    return &RealtimeHandler{
        emailService: emailService,
        logger:      logger,
        metrics:     metrics,
    }
}