package mongodb

import (
	"github.com/bezata/blockchainml-email/internal/monitoring/metrics"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Repository struct {
	db      *mongo.Database
	logger  *zap.Logger
	metrics *metrics.Metrics
	email   *EmailRepository
	staff   *StaffRepository
	thread  *ThreadRepository
}

func NewRepository(db *mongo.Database, logger *zap.Logger, metrics *metrics.Metrics) *Repository {
	return &Repository{
		db:      db,
		logger:  logger,
		metrics: metrics,
		email:   NewEmailRepository(db, logger, metrics),
		staff:   NewStaffRepository(db, logger, metrics),
		thread:  NewThreadRepository(db, logger, metrics),
	}
}

// Implement storage.Repository interface
func (r *Repository) Email() EmailRepository {
	return r.email
}

func (r *Repository) Staff() StaffRepository {
	return r.staff
}

func (r *Repository) Thread() ThreadRepository {
	return r.thread
}
