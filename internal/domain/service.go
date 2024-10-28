package domain

import (
	"context"

	"github.com/bezata/blockchainml-mail/internal/domain/email"
	"github.com/bezata/blockchainml-mail/internal/domain/staff"
	"github.com/bezata/blockchainml-mail/internal/domain/thread"
)

type EmailService interface {
    SendEmail(ctx context.Context, params email.SendEmailParams) (*email.Email, error)
    GetEmail(ctx context.Context, id string) (*email.Email, error)
    ListEmails(ctx context.Context, query email.ListEmailsQuery) ([]*email.Email, error)
    UpdateEmail(ctx context.Context, id string, params email.UpdateEmailParams) (*email.Email, error)
    DeleteEmail(ctx context.Context, id string) error
}

type StaffService interface {
    GetStaff(ctx context.Context, id string) (*staff.Staff, error)
    UpdateStaff(ctx context.Context, id string, params staff.UpdateStaffParams) (*staff.Staff, error)
    ListStaff(ctx context.Context, query staff.ListStaffQuery) ([]*staff.Staff, error)
}

type ThreadService interface {
    GetThread(ctx context.Context, id string) (*thread.Thread, error)
    ListThreads(ctx context.Context, query thread.ListThreadsQuery) ([]*thread.Thread, error)
    UpdateThread(ctx context.Context, id string, params thread.UpdateThreadParams) (*thread.Thread, error)
}
