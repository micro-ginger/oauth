package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

func (repo *repo) Count(q query.Query) (uint64, errors.Error) {
	q = query.NewModelsQuery(q).
		WithModelsHandlerFunc(func() any {
			return new([]scope.Scope)
		})
	return repo.Repository.Count(q)
}

func (repo *repo) ReadItems(q query.Query) ([]*scope.Scope, errors.Error) {
	q = query.NewModelsQuery(q).
		WithModelsHandlerFunc(func() any {
			return new([]scope.Scope)
		})
	items, err := repo.Repository.List(q)
	if err != nil {
		return nil, err
	}
	return *items.(*[]*scope.Scope), nil
}

func (repo *repo) ReadItem(q query.Query) (*scope.Scope, errors.Error) {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(scope.Scope)
		})
	items, err := repo.Repository.Get(q)
	if err != nil {
		return nil, err
	}
	return items.(*scope.Scope), nil
}
