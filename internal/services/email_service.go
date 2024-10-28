package services

import (
	"context"
	"github.com/bezata/blockchainml-email/internal/domain/email"
	"github.com/bezata/blockchainml-email/internal/monitoring/metrics"
	"github.com/bezata/blockchainml-email/pkg/cache"
	"github.com/bezata/blockchainml-email/pkg/search"
	"go.uber.org/zap"
)

type EmailService struct {
	repo      email.Repository
	cache     *cache.Cache
	search    *search.SearchEngine
	logger    *zap.Logger
	metrics   *metrics.Metrics
}

func NewEmailService(
	repo email.Repository,
	cache *cache.Cache,
	search *search.SearchEngine,
	logger *zap.Logger,
	metrics *metrics.Metrics,
) *EmailService {
	return &EmailService{
		repo:    repo,
		cache:   cache,
		search:  search,
		logger:  logger,
		metrics: metrics,
	}
}

type SendEmailParams struct {
	From        string
	To          []string
	Subject     string
	Content     email.EmailContent
	Attachments []email.AttachmentInput
	ThreadID    *string
}

func (s *EmailService) SendEmail(ctx context.Context, params SendEmailParams) (*email.Email, error) {
	// implementation of SendEmail method
}
