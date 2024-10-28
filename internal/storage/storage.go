package storage

import (
	"context"

	"github.com/bezata/blockchainml-email/internal/domain/email"
	"github.com/bezata/blockchainml-email/internal/domain/staff"
	"github.com/bezata/blockchainml-email/internal/domain/thread"
)

// EmailRepository defines email storage operations
type EmailRepository interface {
    Create(ctx context.Context, email *email.Email) error
    Get(ctx context.Context, id string) (*email.Email, error)
    Update(ctx context.Context, email *email.Email) error
    Delete(ctx context.Context, id string) error
    List(ctx context.Context, query *email.ListQuery) ([]*email.Email, error)
}

// StaffRepository defines staff storage operations
type StaffRepository interface {
    Create(ctx context.Context, staff *staff.Staff) error
    Get(ctx context.Context, id string) (*staff.Staff, error)
    GetByEmail(ctx context.Context, email string) (*staff.Staff, error)
    Update(ctx context.Context, staff *staff.Staff) error
    Delete(ctx context.Context, id string) error
    List(ctx context.Context, query *staff.ListQuery) ([]*staff.Staff, error)
}

// ThreadRepository defines thread storage operations
type ThreadRepository interface {
    Create(ctx context.Context, thread *thread.Thread) error
    Get(ctx context.Context, id string) (*thread.Thread, error)
    Update(ctx context.Context, thread *thread.Thread) error
    Delete(ctx context.Context, id string) error
    List(ctx context.Context, query *thread.ListQuery) ([]*thread.Thread, error)
}
