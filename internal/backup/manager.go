package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/bezata/blockchainml-mail/internal/backup"
	"github.com/bezata/blockchainml-mail/internal/config"
	"github.com/bezata/blockchainml-mail/internal/r2"
	"github.com/bezata/blockchainml-mail/internal/storage"
	"github.com/bezata/blockchainml-mail/internal/zap"
	"github.com/google/uuid"
)

type BackupManager struct {
	storage  storage.BackupRepository
	r2Client *r2.Client
	logger   *zap.Logger
	config   *config.Config
}

func (b *BackupManager) CreateBackup(ctx context.Context) error {
	backup := &backup.Backup{
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Status:    "in_progress",
	}

	// Create backup directory
	backupDir := fmt.Sprintf("backups/%s", backup.ID)

	// Backup MongoDB
	mongoBackup, err := b.backupMongoDB(ctx)
	if err != nil {
		return fmt.Errorf("mongodb backup failed: %w", err)
	}

	// Upload MongoDB backup to R2
	mongoKey := fmt.Sprintf("%s/mongodb.gz", backupDir)
	if err := b.r2Client.Upload(ctx, mongoKey, mongoBackup, "application/gzip"); err != nil {
		return fmt.Errorf("failed to upload mongodb backup: %w", err)
	}

	// Backup attachments
	if err := b.backupAttachments(ctx, backupDir); err != nil {
		return fmt.Errorf("attachments backup failed: %w", err)
	}

	backup.Status = "completed"
	backup.R2Path = backupDir
	
	return b.storage.UpdateBackup(ctx, backup)
}
