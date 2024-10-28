package staff

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Staff struct {
    ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Email         string            `bson:"email" json:"email"`
    FullName      string            `bson:"fullName" json:"fullName"`
    Role          string            `bson:"role" json:"role"`
    Department    string            `bson:"department" json:"department"`
    ProfilePhoto  ProfilePhoto      `bson:"profilePhoto" json:"profilePhoto"`
    Status        string            `bson:"status" json:"status"`
    LastActive    time.Time         `bson:"lastActive" json:"lastActive"`
    CreatedAt     time.Time         `bson:"createdAt" json:"createdAt"`
    UpdatedAt     time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type ProfilePhoto struct {
    URL         string    `bson:"url" json:"url"`
    R2Key       string    `bson:"r2Key" json:"r2Key"`
    LastUpdated time.Time `bson:"lastUpdated" json:"lastUpdated"`
}
