package internal

import (
	"context"
	"internal/metrics"
	"internal/r2"
	"internal/storage"

	"go.uber.org/zap"
)

type AuditLogger struct {
	logger     *zap.Logger
	storage    storage.AuditRepository
	r2Client   *r2.Client
	metrics    *metrics.Metrics
}

func (a *AuditLogger) Log(ctx context.Context, event AuditEvent) error {
	// Enrich event with Cloudflare data
	if cfRay := ctx.Value("cf-ray"); cfRay != nil {
		event.Details["cf_ray"] = cfRay.(string)
	}
	if cfCountry := ctx.Value("cf-ipcountry"); cfCountry != nil {
		event.Details["cf_country"] = cfCountry.(string)
	}

	// Store audit event
	if err := a.storage.StoreAuditEvent(ctx, event); err != nil {
		a.logger.Error("failed to store audit event",
			zap.Error(err),
			zap.Any("event", event),
		)
		return err
	}

	// Log high-risk events to R2 for compliance
	if event.Risk == "high" {
		if err := a.archiveToR2(ctx, event); err != nil {
			a.logger.Error("failed to archive high-risk event",
				zap.Error(err),
				zap.Any("event", event),
			)
		}
	}

	return nil
}
