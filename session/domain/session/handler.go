package session

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
)

type SessionHandlerFunc[AccountDetail gateway.ResultGetter] func(ctx context.Context,
	session *Session[AccountDetail]) errors.Error
