package handler

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/authentication/info"
)

type AccountGetter[T account.Model] func(ctx context.Context,
	inf *info.Info[T]) (*account.Account[T], errors.Error)
