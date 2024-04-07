package domain

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

type Repository interface {
	Count(q query.Query) (uint64, errors.Error)
	ReadItems(q query.Query) ([]*scope.Scope, errors.Error)
	ReadItem(q query.Query) (*scope.Scope, errors.Error)
}
