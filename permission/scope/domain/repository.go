package domain

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

type Repository interface {
	Create(q query.Query, scope *scope.Scope) errors.Error
	CreateBulk(q query.Query, scopes []*scope.Scope) errors.Error
	Count(q query.Query) (uint64, errors.Error)
	ReadItems(q query.Query) ([]*scope.Scope, errors.Error)
	ReadItem(q query.Query) (*scope.Scope, errors.Error)
	Delete(q query.Query, name string) errors.Error
	DeleteBulk(q query.Query, names []string) errors.Error
}
