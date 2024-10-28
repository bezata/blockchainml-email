package email

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Participant struct {
	Email        string `bson:"email" json:"email"`
	FullName     string `bson:"fullName" json:"fullName"`
	ProfilePhoto string `bson:"profilePhoto" json:"profilePhoto"`
}

type EmailContent struct {
	Text string `bson:"text" json:"text"`
	HTML string `bson:"html" json:"html"`
}

type Attachment struct {
	Filename    string    `bson:"filename" json:"filename"`
	R2Key       string    `bson:"r2Key" json:"r2Key"`
	ContentType string    `bson:"contentType" json:"contentType"`
	Size        int64     `bson:"size" json:"size"`
	UploadedAt  time.Time `bson:"uploadedAt" json:"uploadedAt"`
}

type AttachmentInput struct {
	Filename    string
	Content     []byte
	ContentType string
}

type EmailFlags struct {
	IsRead      bool `bson:"isRead" json:"isRead"`
	IsStarred   bool `bson:"isStarred" json:"isStarred"`
	IsScheduled bool `bson:"isScheduled" json:"isScheduled"`
	IsDraft     bool `bson:"isDraft" json:"isDraft"`
}

type ThreadInfo struct {
	Depth    int      `bson:"depth" json:"depth"`
	RootID   string   `bson:"rootId" json:"rootId"`
	Path     []string `bson:"path" json:"path"`
}

type EmailMetadata struct {
	ScheduledFor *time.Time `bson:"scheduledFor,omitempty" json:"scheduledFor,omitempty"`
	ClientIP     string     `bson:"clientIp" json:"clientIp"`
	UserAgent    string     `bson:"userAgent" json:"userAgent"`
}

type Email struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MessageID   string            `bson:"messageId" json:"messageId"`
	ThreadID    *string           `bson:"threadId,omitempty" json:"threadId,omitempty"`
	From        Participant       `bson:"from" json:"from"`
	To          []Participant     `bson:"to" json:"to"`
	CC          []Participant     `bson:"cc,omitempty" json:"cc,omitempty"`
	BCC         []Participant     `bson:"bcc,omitempty" json:"bcc,omitempty"`
	Subject     string            `bson:"subject" json:"subject"`
	Content     EmailContent      `bson:"content" json:"content"`
	Attachments []Attachment      `bson:"attachments" json:"attachments"`
	Labels      []string          `bson:"labels" json:"labels"`
	Flags       EmailFlags        `bson:"flags" json:"flags"`
	ThreadInfo  ThreadInfo        `bson:"threadInfo" json:"threadInfo"`
	Metadata    EmailMetadata     `bson:"metadata" json:"metadata"`
	CreatedAt   time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time         `bson:"updatedAt" json:"updatedAt"`
}
