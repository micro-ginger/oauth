package domain

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/accountscope/domain/accountscope"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

type Repository interface {
	Create(q query.Query, item *accountscope.AccountScope) errors.Error
	CreateBulk(ctx context.Context,
		accountId uint64, scopes accountscope.CreateScopeBulk) errors.Error

	GetAllAccountScopes(ctx context.Context,
		accountId uint64, getAll bool) ([]*scope.Detailed, errors.Error)
	GetAccountScopes(ctx context.Context,
		accountId uint64, getAll bool) ([]*scope.Detailed, errors.Error)

	GetAccountScopesFromRoles(ctx context.Context,
		accountId uint64, roles []string, getAll bool) ([]*scope.Detailed, errors.Error)
	GetAccountScopesFromScopes(ctx context.Context,
		accountId uint64, names []string, getAll bool) ([]*scope.Detailed, errors.Error)

	ListDefaultAccountScopes(ctx context.Context,
		accountId uint64, getAll bool) ([]*scope.Detailed, errors.Error)
	ListUnauthorizedAccountScopes(ctx context.Context,
		accountId uint64) ([]*scope.Scope, errors.Error)

	Delete(ctx context.Context, accountId uint64, scopeId uint64) errors.Error
	DeleteBulk(ctx context.Context,
		accountId uint64, scopeIds []uint64) errors.Error
}
