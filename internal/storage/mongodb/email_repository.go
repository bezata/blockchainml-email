package mongodb

import (
	"context"
	"time"

	"github.com/bezata/blockchainml-mail/internal/monitoring/metrics"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type EmailRepository struct {
    db         *mongo.Database
    collection *mongo.Collection
    logger     *zap.Logger
    metrics    *metrics.Metrics
}

func NewEmailRepository(db *mongo.Database, logger *zap.Logger, metrics *metrics.Metrics) *EmailRepository {
    return &EmailRepository{
        db:         db,
        collection: db.Collection("emails"),
        logger:     logger,
        metrics:    metrics,
    }
}

// Advanced query builder for flexible email searching
type EmailQuery struct {
    ThreadID     *string
    Labels       []string
    SearchText   string
    DateRange    *DateRange
    Participants []string
    HasFlags     *Flags
    Sort         []string
    Page         int64
    PageSize     int64
}

func (r *EmailRepository) FindEmails(ctx context.Context, query EmailQuery) ([]*Email, error) {
    startTime := time.Now()
    defer func() {
        r.metrics.DatabaseLatency.WithLabelValues("find_emails").Observe(time.Since(startTime).Seconds())
    }()

    filter := bson.M{}
    
    // Build complex query
    if query.ThreadID != nil {
        filter["threadId"] = *query.ThreadID
    }

    if len(query.Labels) > 0 {
        filter["labels"] = bson.M{"$in": query.Labels}
    }

    if len(query.Participants) > 0 {
        filter["$or"] = bson.A{
            bson.M{"from.email": bson.M{"$in": query.Participants}},
            bson.M{"to.email": bson.M{"$in": query.Participants}},
            bson.M{"cc.email": bson.M{"$in": query.Participants}},
        }
    }

    if query.SearchText != "" {
        filter["$text"] = bson.M{"$search": query.SearchText}
    }

    if query.DateRange != nil {
        dateFilter := bson.M{}
        if !query.DateRange.Start.IsZero() {
            dateFilter["$gte"] = query.DateRange.Start
        }
        if !query.DateRange.End.IsZero() {
            dateFilter["$lte"] = query.DateRange.End
        }
        filter["createdAt"] = dateFilter
    }

    // Apply flag filters
    if query.HasFlags != nil {
        if query.HasFlags.IsRead != nil {
            filter["flags.isRead"] = *query.HasFlags.IsRead
        }
        if query.HasFlags.IsStarred != nil {
            filter["flags.isStarred"] = *query.HasFlags.IsStarred
        }
    }

    // Configure options
    opts := options.Find().
        SetLimit(query.PageSize).
        SetSkip((query.Page - 1) * query.PageSize)

    if len(query.Sort) > 0 {
        sort := bson.D{}
        for _, s := range query.Sort {
            direction := 1
            if s[0] == '-' {
                direction = -1
                s = s[1:]
            }
            sort = append(sort, bson.E{Key: s, Value: direction})
        }
        opts.SetSort(sort)
    }

    // Execute query with timeout
    queryCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    cursor, err := r.collection.Find(queryCtx, filter, opts)
    if err != nil {
        r.logger.Error("failed to find emails", zap.Error(err))
        return nil, err
    }
    defer cursor.Close(ctx)

    var emails []*Email
    if err = cursor.All(ctx, &emails); err != nil {
        r.logger.Error("failed to decode emails", zap.Error(err))
        return nil, err
    }

    return emails, nil
}
