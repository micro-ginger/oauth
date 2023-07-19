package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/register/domain/register"
)

func (repo *repo[T]) Begin(q query.Query,
	options ...repository.Options) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(register.Register[T])
		})
	return repo.base.Begin(q, options...)
}

func (repo *repo[T]) End(q query.Query) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(register.Register[T])
		})
	return repo.base.End(q)
}
