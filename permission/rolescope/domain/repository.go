package domain

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/rolescope/domain/rolescope"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

type Repository interface {
	Create(q query.Query, item *rolescope.RoleScope) errors.Error
	CreateBulk(ctx context.Context, roleId uint64, scopeIds []uint64) errors.Error
	CountScopes(q query.Query) (uint64, errors.Error)
	ListScopes(q query.Query) ([]*scope.Scope, errors.Error)
	Delete(ctx context.Context, roleId uint64, scopeId uint64) errors.Error
	DeleteBulk(ctx context.Context, roleId uint64, scopeIds []uint64) errors.Error
}
