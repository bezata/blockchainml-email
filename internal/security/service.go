package security

import (
	"context"

	"github.com/bezata/blockchainml-mail/internal/config"
	"go.uber.org/zap"
)

type Service struct {
    config     *config.SecurityConfig
    logger     *zap.Logger
    encryptor  *Encryptor
    audit      *AuditLogger
    rateLimit  *RateLimiter
}

func NewService(cfg *config.SecurityConfig, logger *zap.Logger) *Service {
    encryptor, err := NewEncryptor([]byte(cfg.EncryptionKey))
    if err != nil {
        logger.Fatal("failed to initialize encryptor", zap.Error(err))
    }

    return &Service{
        config:    cfg,
        logger:    logger,
        encryptor: encryptor,
        audit:     NewAuditLogger(logger),
        rateLimit: NewRateLimiter(cfg.RateLimit),
    }
}

// ValidateCloudflareToken validates the Cloudflare Workers token
func (s *Service) ValidateCloudflareToken(ctx context.Context, token string) (bool, error) {
    // Implement Cloudflare Workers token validation
    return true, nil
}

// EncryptForCloudflare encrypts data for Cloudflare Workers
func (s *Service) EncryptForCloudflare(data []byte) (string, error) {
    return s.encryptor.Encrypt(data)
}

// DecryptFromCloudflare decrypts data from Cloudflare Workers
func (s *Service) DecryptFromCloudflare(data string) ([]byte, error) {
    return s.encryptor.Decrypt(data)
}
