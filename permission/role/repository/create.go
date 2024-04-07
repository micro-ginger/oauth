package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/role/domain/role"
)

func (repo *repo) Create(q query.Query, item *role.Role) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(role.Role)
		})
	err := repo.Repository.Create(q, item)
	if err != nil {
		return err
	}
	return nil
}
