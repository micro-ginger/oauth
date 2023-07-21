package session

import (
	"context"

	"github.com/ginger-core/errors"
)

type SessionHandlerFunc func(ctx context.Context, session *Session) errors.Error
