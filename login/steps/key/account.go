package key

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type GetAccountHandlerFunc[T account.Model] func(ctx context.Context,
	key string) (*account.Account[T], errors.Error)
