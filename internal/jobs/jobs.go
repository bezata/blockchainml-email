package jobs

import "time"

const (
	TaskSendScheduledEmail     = "send_scheduled_email"
	TaskProcessAttachments     = "process_attachments"
	TaskUpdateSearchIndex      = "update_search_index"
	TaskGenerateEmailAnalytics = "generate_email_analytics"
)

type ScheduledEmailPayload struct {
	EmailID string `json:"emailId"`
}

type AttachmentProcessingPayload struct {
	EmailID      string   `json:"emailId"`
	AttachmentID string   `json:"attachmentId"`
	Operations   []string `json:"operations"`
}

type SearchIndexPayload struct {
	EmailID string `json:"emailId"`
	Action  string `json:"action"` // index, update, delete
}

type EmailAnalyticsPayload struct {
	EmailID   string    `json:"emailId"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
