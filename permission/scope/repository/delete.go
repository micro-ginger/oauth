package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

func (repo *repo) Delete(q query.Query, name string) errors.Error {
	q = query.NewFilter(q).
		WithMatch(&query.Match{
			Key:      "name",
			Operator: query.Equal,
			Value:    name,
		})
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(scope.Scope)
		})
	if err := repo.Repository.Delete(q); err != nil {
		return err
	}
	return nil
}

func (repo *repo) DeleteBulk(q query.Query, names []string) errors.Error {
	q = query.NewFilter(q).
		WithMatch(&query.Match{
			Key:      "name",
			Operator: query.In,
			Value:    names,
		})
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(scope.Scope)
		})
	if err := repo.Repository.Delete(q); err != nil {
		return err
	}
	return nil
}
