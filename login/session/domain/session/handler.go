package session

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type Handler[acc account.Model] interface {
	Generate(ctx context.Context,
		request *GenerateRequest) (*Session[acc], errors.Error)

	Save(ctx context.Context, info *Session[acc]) errors.Error

	Get(ctx context.Context, challenge string) (*Session[acc], errors.Error)

	Delete(ctx context.Context, challenge string) errors.Error
}
