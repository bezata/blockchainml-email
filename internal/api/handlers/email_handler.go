package handlers

import (
    "net/http"
    "encoding/json"
    "github.com/bezata/blockchainml-email/internal/services"  // Add this import 
)

type SendEmailRequest struct {
    To          []string          `json:"to" validate:"required,min=1,dive,email"`
    Subject     string            `json:"subject" validate:"required"`
    Content     EmailContent      `json:"content" validate:"required"`
    Attachments []AttachmentInput `json:"attachments,omitempty"`
    ThreadID    *string           `json:"threadId,omitempty"`
}

type EmailContent struct {
    Text string `json:"text" validate:"required"`
    HTML string `json:"html,omitempty"`
}

type AttachmentInput struct {
    Filename    string `json:"filename" validate:"required"`
    ContentType string `json:"contentType" validate:"required"`
    Content     []byte `json:"content" validate:"required"`
}

func (h *EmailHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    
    var req SendEmailRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        h.respondError(w, http.StatusBadRequest, "Invalid request body")
        return
    }

    email, err := h.emailService.SendEmail(ctx, services.SendEmailParams{
        From:        ctx.Value("userId").(string),
        To:          req.To,
        Subject:     req.Subject,
        Content:     req.Content,
        Attachments: req.Attachments,
        ThreadID:    req.ThreadID,
    })
    if err != nil {
        h.logger.Error("failed to send email", zap.Error(err))
        h.respondError(w, http.StatusInternalServerError, "Failed to send email")
        return
    }

    h.respondJSON(w, http.StatusCreated, email)
}

func (h *EmailHandler) respondJSON(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func (h *EmailHandler) respondError(w http.ResponseWriter, status int, message string) {
    h.respondJSON(w, status, map[string]string{"error": message})
}
