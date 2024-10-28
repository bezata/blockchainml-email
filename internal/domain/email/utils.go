package email

import (
	"github.com/google/uuid"
)

func generateMessageID() string {
	return uuid.New().String()
}

func generateThreadID() string {
	return uuid.New().String()
}

func truncateText(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
