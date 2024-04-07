package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/role/domain/role"
)

func (repo *repo) Update(q query.Query, item *role.UpdateRequest) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(role.Role)
		})
	if err := repo.Repository.Update(q, nil); err != nil {
		return err
	}
	return nil
}
