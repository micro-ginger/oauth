package account

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type UseCase[T a.ExtentedModel] interface {
	Update(ctx context.Context, q query.Query,
		update *account.Update[T]) errors.Error
}
