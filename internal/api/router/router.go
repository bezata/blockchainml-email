package router

import (
    "github.com/gin-gonic/gin"
    "github.com/bezata/blockchainml-email/internal/api/handlers"
    "github.com/bezata/blockchainml-email/internal/api/middleware"
)

func NewRouter(handlers *handlers.Handlers, mw *middleware.Middleware) *gin.Engine {
    router := gin.Default()

    // Add middleware
    router.Use(mw.Logger.Handle())
    router.Use(mw.Metrics.Handle())
    router.Use(mw.RateLimit.Handle())

    // API routes
    api := router.Group("/api/v1")
    {
        // Public routes
        api.POST("/auth/login", handlers.Auth.Login)
        api.POST("/auth/refresh", handlers.Auth.RefreshToken)

        // Protected routes
        protected := api.Group("")
        protected.Use(mw.Auth.Handle())
        {
            // Email routes
            protected.POST("/emails", handlers.Email.SendEmail)
            protected.GET("/emails/search", handlers.Email.SearchEmails)
            // Add other routes...
        }
    }

    return router
}
