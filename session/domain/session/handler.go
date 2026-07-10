package session

import (
	"context"

	"github.com/ginger-core/errors"
	a "github.com/micro-blonde/auth/account"
)

type SessionHandlerFunc[AccountDetail a.ExtentedModel] func(ctx context.Context,
	session *Session[AccountDetail]) errors.Error
