package domain

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

type UseCase interface {
	Create(ctx context.Context, scope *scope.Scope) errors.Error
	CreateBulk(ctx context.Context, scopes []*scope.Scope) errors.Error
	Count(ctx context.Context, q query.Query) (uint64, errors.Error)
	Paginate(ctx context.Context, q query.Query) ([]*scope.Scope, errors.Error)
	Get(ctx context.Context, id uint64) (*scope.Scope, errors.Error)
	Delete(ctx context.Context, name string) errors.Error
	DeleteBulk(ctx context.Context, names []string) errors.Error
}
