package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/account/profile/domain/profile"
)

func (repo *repo[T]) Upsert(q query.Query,
	entity *profile.Profile[T]) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(profile.Profile[T])
		})
	if err := repo.Repository.Upsert(q, entity); err != nil {
		return err
	}
	return nil
}
