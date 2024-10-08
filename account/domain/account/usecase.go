package account

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/account"
)

type UseCase[T account.Model] interface {
	Create(ctx context.Context, account *Account[T]) errors.Error

	Count(ctx context.Context, q query.Query) (uint64, errors.Error)
	List(ctx context.Context, q query.Query) ([]*Account[T], errors.Error)
	Get(ctx context.Context, q query.Query) (*Account[T], errors.Error)
	GetById(ctx context.Context, id uint64) (*Account[T], errors.Error)

	Update(ctx context.Context, q query.Query, update *Update[T]) errors.Error
	UpdateAccount(ctx context.Context,
		q query.Query, update *Account[T]) errors.Error

	Upsert(ctx context.Context, q query.Query,
		account *Account[T]) errors.Error

	VerifyPassword(ctx context.Context,
		account *Account[T], password string) errors.Error

	ValidatePassword(ctx context.Context, password string) errors.Error
	UpdatePassword(ctx context.Context,
		q query.Query, hashedPassword []byte) errors.Error
}

type ResetPasswordHandler interface {
	ValidatePassword(ctx context.Context, password string) errors.Error
	UpdatePassword(ctx context.Context,
		q query.Query, hashedPassword []byte) errors.Error
}

type PasswordUpdateHandler[T account.Model] interface {
	Get(ctx context.Context, q query.Query) (*Account[T], errors.Error)
	ValidatePassword(ctx context.Context, password string) errors.Error
	UpdatePassword(ctx context.Context,
		q query.Query, hashedPassword []byte) errors.Error
}
