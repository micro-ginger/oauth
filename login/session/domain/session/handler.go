package session

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type Handler[acc account.Model, SessionAccountDetail gateway.ResultGetter] interface {
	Generate(
		ctx context.Context, request *GenerateRequest,
	) (*Session[acc, SessionAccountDetail], errors.Error)

	Save(ctx context.Context, info *Session[acc, SessionAccountDetail]) errors.Error

	Get(
		ctx context.Context, challenge string,
	) (*Session[acc, SessionAccountDetail], errors.Error)

	Delete(ctx context.Context, challenge string) errors.Error
}
