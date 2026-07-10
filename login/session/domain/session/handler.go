package session

import (
	"context"

	"github.com/ginger-core/errors"
	a "github.com/micro-blonde/auth/account"
)

type Handler[extendedAcc a.ExtentedModel] interface {
	Generate(
		ctx context.Context, request *GenerateRequest,
	) (*Session[extendedAcc], errors.Error)

	Save(ctx context.Context, info *Session[extendedAcc]) errors.Error

	Get(
		ctx context.Context, challenge string,
	) (*Session[extendedAcc], errors.Error)

	Delete(ctx context.Context, challenge string) errors.Error
}
