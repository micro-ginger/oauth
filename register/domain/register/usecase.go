package register

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type UseCase[T Model, acc account.Model] interface {
	Base() UseCase[T, acc]
	Wrap(uc UseCase[T, acc])

	Register(ctx context.Context, request *Request[T, acc]) errors.Error
}
