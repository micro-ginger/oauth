package info

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type AccountGetter[T account.Model] func(ctx context.Context,
	inf *Info[T]) (*account.Account[T], errors.Error)
