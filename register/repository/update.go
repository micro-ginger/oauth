package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/register/domain/register"
)

func (repo *repo[T]) Update(q query.Query,
	update *register.Register[T]) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(register.Register[T])
		})
	if err := repo.base.Update(q, update); err != nil {
		return err
	}
	return nil
}
