package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/role/domain/role"
)

func (repo *repo) Count(q query.Query) (uint64, errors.Error) {
	q = query.NewModelsQuery(q).
		WithModelsHandlerFunc(func() any {
			return new([]role.Role)
		})
	return repo.Repository.Count(q)
}

func (repo *repo) List(q query.Query) ([]*role.Role, errors.Error) {
	q = query.NewModelsQuery(q).
		WithModelsHandlerFunc(func() any {
			return new([]role.Role)
		})
	items, err := repo.Repository.List(q)
	if err != nil {
		return nil, err
	}
	return *items.(*[]*role.Role), nil
}

func (repo *repo) Get(q query.Query) (*role.Role, errors.Error) {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(role.Role)
		})
	items, err := repo.Repository.Get(q)
	if err != nil {
		return nil, err
	}
	return items.(*role.Role), nil
}
