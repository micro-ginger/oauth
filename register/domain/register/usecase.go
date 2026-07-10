package register

import (
	"context"

	"github.com/ginger-core/errors"
	a "github.com/micro-blonde/auth/account"
)

type UseCase[T Model, acc a.ExtentedModel] interface {
	Base() UseCase[T, acc]
	Wrap(uc UseCase[T, acc])

	Register(ctx context.Context, request *Request[T, acc]) errors.Error
}
