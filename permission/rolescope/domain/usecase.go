package domain

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-repository/sql/query"
	"github.com/micro-ginger/oauth/permission/rolescope/domain/rolescope"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

type UseCase interface {
	Create(ctx context.Context, item *rolescope.RoleScope) errors.Error
	CreateBulk(ctx context.Context, roleId uint64, scopeIds []uint64) errors.Error
	CountScopes(ctx context.Context, q query.Query) (uint64, errors.Error)
	ListScopes(ctx context.Context, q query.Query) ([]*scope.Scope, errors.Error)
	Delete(ctx context.Context, roleId uint64, scopeId uint64) errors.Error
	DeleteBulk(ctx context.Context, roleId uint64, scopeIds []uint64) errors.Error
}
