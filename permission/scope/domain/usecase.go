package domain

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

type UseCase interface {
	Count(ctx context.Context, q query.Query) (uint64, errors.Error)
	Paginate(ctx context.Context, q query.Query) ([]*scope.Scope, errors.Error)
	Get(ctx context.Context, id uint64) (*scope.Scope, errors.Error)
}
