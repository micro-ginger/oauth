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
	Set(ctx context.Context, challenge, key string, value any) errors.Error

	Get(ctx context.Context, challenge string) (*Session[acc], errors.Error)
	GetItem(ctx context.Context, challenge, key string, ref any) errors.Error

	Delete(ctx context.Context, challenge string) errors.Error
	DeleteItem(ctx context.Context, challenge, key string) errors.Error
}
