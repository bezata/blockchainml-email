package r2

import (
	"context"
	"fmt"
	"time"

	"github.com/bezata/blockchainml-email/internal/domain/email"
	"go.uber.org/zap"
)

type Storage struct {
    client *Client
    logger *zap.Logger
}

func NewStorage(client *Client, logger *zap.Logger) *Storage {
    return &Storage{
        client: client,
        logger: logger,
    }
}

// StoreAttachment stores an email attachment in R2
func (s *Storage) StoreAttachment(ctx context.Context, emailID string, attachment email.AttachmentInput) (*email.Attachment, error) {
    key := fmt.Sprintf("attachments/%s/%s", emailID, attachment.Filename)
    
    if err := s.client.Upload(ctx, key, attachment.Content, attachment.ContentType); err != nil {
        return nil, err
    }

    return &email.Attachment{
        Filename:    attachment.Filename,
        R2Key:      key,
        ContentType: attachment.ContentType,
        Size:       int64(len(attachment.Content)),
        UploadedAt: time.Now(),
    }, nil
}

// GetAttachment retrieves an attachment from R2
func (s *Storage) GetAttachment(ctx context.Context, key string) ([]byte, error) {
    return s.client.Download(ctx, key)
}

// DeleteAttachments deletes all attachments for an email
func (s *Storage) DeleteAttachments(ctx context.Context, emailID string) error {
    prefix := fmt.Sprintf("attachments/%s/", emailID)
    
    keys, err := s.client.ListObjects(ctx, prefix)
    if err != nil {
        return err
    }

    for _, key := range keys {
        if err := s.client.Delete(ctx, key); err != nil {
            s.logger.Error("failed to delete attachment",
                zap.String("key", key),
                zap.Error(err),
            )
            // Continue deleting other attachments
            continue
        }
    }

    return nil
}
