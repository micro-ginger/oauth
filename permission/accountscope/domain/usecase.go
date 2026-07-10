package domain

import (
	"context"

	"github.com/ginger-core/errors"
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/permission/accountscope/domain/accountscope"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type UseCase[SessionAccountDetail a.ExtentedModel] interface {
	RegisterCreateEventHandle(handle accountscope.CreatedScopeEventHandle)

	SessionAddRequestedRoleScopes(ctx context.Context,
		session *session.Session[SessionAccountDetail]) errors.Error
	SessionRemoveUnauthorized(ctx context.Context,
		session *session.Session[SessionAccountDetail]) errors.Error

	Create(ctx context.Context, item *accountscope.AccountScope) errors.Error
	CreateBulk(ctx context.Context, accountId uint64,
		scopes accountscope.CreateScopeBulk) errors.Error

	GetAllAccountScopes(ctx context.Context,
		accountId uint64, getAll bool) ([]*scope.Detailed, errors.Error)
	GetAccountScopes(ctx context.Context,
		accountId uint64, getAll bool) ([]*scope.Detailed, errors.Error)
	GetAccountScopesFromRoles(ctx context.Context,
		accountId uint64, roles []string, getAll bool) ([]*scope.Detailed, errors.Error)
	ListDefaultAccountScopes(ctx context.Context,
		accountId uint64, getAll bool) ([]*scope.Detailed, errors.Error)

	Delete(ctx context.Context, accountId uint64, scopeId uint64) errors.Error
	DeleteBulk(ctx context.Context, accountId uint64, scopeIds []uint64) errors.Error

	Authorize(ctx context.Context, accountId uint64, scopes ...string) errors.Error
	Revoke(ctx context.Context, accountId uint64, scopes ...string) errors.Error
}
