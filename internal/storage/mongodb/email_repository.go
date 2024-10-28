package mongodb

import (
	"context"
	"time"

	"github.com/bezata/blockchainml-email/internal/monitoring/metrics"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"github.com/bezata/blockchainml-email/internal/domain/email"
)

type EmailRepository struct {
	collection *mongo.Collection
	logger     *zap.Logger
	metrics    *metrics.Metrics
}

func NewEmailRepository(db *mongo.Database, logger *zap.Logger, metrics *metrics.Metrics) *EmailRepository {
	return &EmailRepository{
		collection: db.Collection("emails"),
		logger:     logger,
		metrics:    metrics,
	}
}

func (r *EmailRepository) Create(ctx context.Context, email *email.Email) error {
	startTime := time.Now()
	defer func() {
		r.metrics.DatabaseLatency.WithLabelValues("create_email").Observe(time.Since(startTime).Seconds())
	}()

	_, err := r.collection.InsertOne(ctx, email)
	if err != nil {
		r.logger.Error("failed to create email", zap.Error(err))
		return err
	}

	return nil
}

func (r *EmailRepository) Get(ctx context.Context, id string) (*email.Email, error) {
	startTime := time.Now()
	defer func() {
		r.metrics.DatabaseLatency.WithLabelValues("get_email").Observe(time.Since(startTime).Seconds())
	}()

	var result email.Email
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		r.logger.Error("failed to get email", zap.Error(err))
		return nil, err
	}

	return &result, nil
}
