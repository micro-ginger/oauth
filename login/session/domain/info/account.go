package info

import (
	"context"

	"github.com/ginger-core/errors"
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type AccountGetter[T a.ExtentedModel] func(ctx context.Context,
	inf *Info[T]) (*account.Account[T], errors.Error)
