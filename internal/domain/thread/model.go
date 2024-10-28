package thread

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Thread struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    ThreadID     string            `bson:"threadId" json:"threadId"`
    Subject      string            `bson:"subject" json:"subject"`
    Participants []Participant     `bson:"participants" json:"participants"`
    LastMessage  LastMessage       `bson:"lastMessage" json:"lastMessage"`
    MessageCount int               `bson:"messageCount" json:"messageCount"`
    CreatedAt    time.Time         `bson:"createdAt" json:"createdAt"`
    UpdatedAt    time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type Participant struct {
    Email        string `bson:"email" json:"email"`
    FullName     string `bson:"fullName" json:"fullName"`
    ProfilePhoto string `bson:"profilePhoto" json:"profilePhoto"`
}

type LastMessage struct {
    MessageID string    `bson:"messageId" json:"messageId"`
    From      string    `bson:"from" json:"from"`
    Subject   string    `bson:"subject" json:"subject"`
    Snippet   string    `bson:"snippet" json:"snippet"`
    SentAt    time.Time `bson:"sentAt" json:"sentAt"`
}
