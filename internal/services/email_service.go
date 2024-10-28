package services

import (
	"github.com/bezata/blockchainml-mail/internal/domain/email"
	"github.com/bezata/blockchainml-mail/pkg/cache"
	"github.com/bezata/blockchainml-mail/pkg/search"
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
