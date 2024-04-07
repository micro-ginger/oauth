package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
	"gorm.io/gorm"
)

func (repo *repo) improveQuery(query query.Query) {
	db := repo.GetDB(query).(*gorm.DB).
		Table("scopes").
		Select("id, created_at, updated_at, name, state, description").
		Joins("JOIN role_scopes ON scope_id=id")
	query.SetDB(db)
}

func (repo *repo) CountScopes(q query.Query) (uint64, errors.Error) {
	repo.improveQuery(q)
	q = query.NewModelsQuery(q).
		WithModelsHandlerFunc(func() any {
			return new([]*scope.Scope)
		})
	return repo.Repository.Count(q)
}

func (repo *repo) ListScopes(q query.Query) ([]*scope.Scope, errors.Error) {
	repo.improveQuery(q)
	q = query.NewModelsQuery(q).
		WithModelsHandlerFunc(func() any {
			return new([]*scope.Scope)
		})
	items, err := repo.Repository.List(q)
	if err != nil {
		return nil, err
	}
	return *items.(*[]*scope.Scope), nil
}
