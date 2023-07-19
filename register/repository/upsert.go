package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/register/domain/register"
)

func (repo *repo[T]) Upsert(q query.Query,
	entity *register.Register[T]) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(register.Register[T])
		})
	if err := repo.base.Upsert(q, entity); err != nil {
		return err
	}
	return nil
}
