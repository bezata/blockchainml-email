package email

import (
	"time"

	"github.com/bezata/blockchainml-email/internal/monitoring/metrics"
	"github.com/bezata/blockchainml-email/pkg/cache"
	"github.com/bezata/blockchainml-email/pkg/search"
	"go.uber.org/zap"
)

type Repository interface {
	SendEmail(params SendEmailParams) error
}

type Service struct {
	repo      Repository
	cache     *cache.Cache
	search    *search.SearchEngine
	logger    *zap.Logger
	metrics   *metrics.Metrics
}

type SendEmailParams struct {
	From        string
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Content     EmailContent
	Attachments []AttachmentInput
	ThreadID    *string
	Schedule    *time.Time
}

func NewService(
	repo Repository,
	cache *cache.Cache,
	search *search.SearchEngine,
	logger *zap.Logger,
	metrics *metrics.Metrics,
) *Service {
	return &Service{
		repo:    repo,
		cache:   cache,
		search:  search,
		logger:  logger,
		metrics: metrics,
	}
}
