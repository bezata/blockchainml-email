package email

import (
	"context"
	"fmt"
	"time"

	"github.com/bezata/blockchainml-mail/internal/monitoring/metrics"
	"github.com/bezata/blockchainml-mail/internal/queue"
	"github.com/bezata/blockchainml-mail/internal/realtime"
	"github.com/bezata/blockchainml-mail/internal/security"
	"github.com/bezata/blockchainml-mail/internal/storage"
	"github.com/bezata/blockchainml-mail/pkg/cache"
	"github.com/bezata/blockchainml-mail/pkg/r2"
	"github.com/bezata/blockchainml-mail/pkg/search"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service struct {
	repo          storage.EmailRepository
	cache         cache.Cache
	search        search.Engine
	queue         queue.Queue
	r2Client      *r2.Client
	logger        *zap.Logger
	metrics       *metrics.Metrics
	notifier      *realtime.Notifier
	security      *security.Service
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

func (s *Service) SendEmail(ctx context.Context, params SendEmailParams) (*Email, error) {
	startTime := time.Now()
	defer func() {
		s.metrics.EmailLatency.WithLabelValues("send").Observe(time.Since(startTime).Seconds())
	}()

	// Create email
	email := &Email{
		MessageID:  uuid.New().String(),
		From:       params.From,
		To:         params.To,
		CC:         params.CC,
		BCC:        params.BCC,
		Subject:    params.Subject,
		Content:    params.Content,
		CreatedAt:  time.Now(),
		Status:     "pending",
	}

	// Process attachments through R2
	if len(params.Attachments) > 0 {
		attachments, err := s.processAttachments(ctx, params.Attachments)
		if err != nil {
			return nil, err
		}
		email.Attachments = attachments
	}

	// Handle scheduled emails
	if params.Schedule != nil {
		return s.scheduleEmail(ctx, email, *params.Schedule)
	}

	// Process thread information
	if err := s.processThreadInfo(ctx, email); err != nil {
		return nil, err
	}

	// Encrypt sensitive content
	encryptedContent, err := s.security.EncryptForCloudflare([]byte(email.Content.Text))
	if err != nil {
		s.logger.Error("failed to encrypt email content", zap.Error(err))
		return nil, err
	}
	email.Content.Text = encryptedContent

	// Save email
	saved, err := s.repo.Create(ctx, email)
	if err != nil {
		s.logger.Error("failed to save email", zap.Error(err))
		return nil, err
	}

	// Update search index
	if err := s.search.IndexEmail(ctx, saved); err != nil {
		s.logger.Error("failed to index email", zap.Error(err))
	}

	// Send real-time notifications
	s.notifyRecipients(ctx, saved)

	return saved, nil
}

// Email scheduling handler
func (s *Service) scheduleEmail(ctx context.Context, email *Email, scheduleTime time.Time) (*Email, error) {
	// Set scheduled flag
	email.Flags.IsScheduled = ptr(true)
	email.Metadata.ScheduledFor = &scheduleTime

	// Save as draft
	email.Labels = append(email.Labels, "scheduled")
	saved, err := s.repo.Create(ctx, email)
	if err != nil {
		return nil, err
	}

	// Add to scheduling queue
	task := queue.Task{
		Type:      "send_scheduled_email",
		Payload:   email.ID.Hex(),
		RunAt:     scheduleTime,
		Priority:  queue.PriorityHigh,
	}

	if err := s.queue.Schedule(ctx, task); err != nil {
		s.logger.Error("failed to schedule email",
			zap.Error(err),
			zap.String("messageId", email.MessageID),
		)
		return nil, err
	}

	return saved, nil
}

// Real-time notification handler
func (s *Service) notifyRecipients(ctx context.Context, email *Email) {
	recipients := make([]string, 0, len(email.To)+len(email.CC))
	recipients = append(recipients, email.To...)
	recipients = append(recipients, email.CC...)

	notification := realtime.Notification{
		Type: "new_email",
		Data: map[string]interface{}{
			"messageId": email.MessageID,
			"from":     email.From,
			"subject":  email.Subject,
			"preview":  truncateText(email.Content.Text, 100),
		},
	}

	for _, recipient := range recipients {
		s.notifier.NotifyUser(ctx, recipient, notification)
	}
}

// Thread processing
func (s *Service) processThreadInfo(ctx context.Context, email *Email) error {
	if email.ThreadID == nil {
		threadID := generateThreadID()
		email.ThreadID = &threadID
		email.ThreadInfo = ThreadInfo{
			Depth:    0,
			RootID:   threadID,
			Path:     []string{threadID},
		}
		return nil
	}

	// Get parent email if this is a reply
	if email.ParentID != nil {
		parent, err := s.repo.FindByID(ctx, *email.ParentID)
		if err != nil {
			return err
		}

		// Update thread info
		email.ThreadInfo = ThreadInfo{
			Depth:    parent.ThreadInfo.Depth + 1,
			RootID:   parent.ThreadInfo.RootID,
			Path:     append(parent.ThreadInfo.Path, email.MessageID),
		}

		// Update parent's reply count
		parent.ThreadInfo.ReplyCount++
		parent.ThreadInfo.LastReplyAt = &email.CreatedAt
		
		if err := s.repo.Update(ctx, parent); err != nil {
			s.logger.Error("failed to update parent email",
				zap.Error(err),
				zap.String("parentId", parent.MessageID),
			)
		}
	}

	return nil
}

// Helper function to create email participant
func (s *Service) createParticipant(email string) Participant {
	// Get staff info from cache or database
	staff, err := s.getStaffByEmail(ctx, email)
	if err != nil {
		return Participant{
			Email: email,
		}
	}

	return Participant{
		Email:        email,
		FullName:     staff.FullName,
		ProfilePhoto: staff.ProfilePhoto.URL,
	}
}

// Attachment processing
func (s *Service) processAttachments(ctx context.Context, inputs []AttachmentInput) ([]Attachment, error) {
	attachments := make([]Attachment, 0, len(inputs))

	for _, input := range inputs {
		// Generate R2 key
		key := fmt.Sprintf("attachments/%s/%s", uuid.New().String(), input.Filename)

		// Upload to R2
		if err := s.r2Client.Upload(ctx, key, input.Content, input.ContentType); err != nil {
			return nil, err
		}

		attachment := Attachment{
			Filename:    input.Filename,
			R2Key:      key,
			ContentType: input.ContentType,
			Size:       int64(len(input.Content)),
			UploadedAt: time.Now(),
		}

		attachments = append(attachments, attachment)
	}

	return attachments, nil
}

// Helper function for creating boolean pointers
func ptr(b bool) *bool {
	return &b
}
