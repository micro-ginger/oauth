package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/register/domain/register"
)

func (repo *repo[T]) Get(q query.Query) (*register.Register[T], errors.Error) {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(register.Register[T])
		})
	r, err := repo.base.Get(q)
	if err != nil {
		return nil, err
	}
	return r.(*register.Register[T]), nil
}
