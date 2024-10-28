package middleware

import (
	"github.com/bezata/blockchainml-mail/internal/security"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CloudflareMiddleware struct {
	security *security.Service
	logger   *zap.Logger
}

func NewCloudflareMiddleware(security *security.Service, logger *zap.Logger) *CloudflareMiddleware {
	return &CloudflareMiddleware{
		security: security,
		logger:   logger,
	}
}

func (m *CloudflareMiddleware) ValidateAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate Cloudflare Access JWT
		token := c.GetHeader("Cf-Access-Jwt-Assertion")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "No Cloudflare Access token"})
			return
		}

		valid, err := m.security.ValidateCloudflareToken(c, token)
		if err != nil || !valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid Cloudflare Access token"})
			return
		}

		c.Next()
	}
}
