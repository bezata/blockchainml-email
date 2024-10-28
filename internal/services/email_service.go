package services

import (
	"github.com/bezata/blockchainml-email/internal/domain/email"
	"github.com/bezata/blockchainml-email/pkg/cache"
	"github.com/bezata/blockchainml-email/pkg/search"
	"go.uber.org/zap"
)

type EmailService struct {
    repo      email.Repository
    cache     *cache.Cache
    search    *search.SearchEngine
    logger    *zap.Logger
}

func NewEmailService(
    repo email.Repository,
    cache *cache.Cache,
    search *search.SearchEngine,
    logger *zap.Logger,
) *EmailService {
    return &EmailService{
        repo:   repo,
        cache:  cache,
        search: search,
        logger: logger,
    }
}
