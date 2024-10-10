package account

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type UseCase[T account.Model] interface {
	Update(ctx context.Context, q query.Query,
		update *account.Update[T]) errors.Error
}
