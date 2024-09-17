package account

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-blonde/auth/account"
)

type Manager[T account.Model] interface {
	HandleInternalStatus(ctx context.Context, account *Account[T]) errors.Error
}
