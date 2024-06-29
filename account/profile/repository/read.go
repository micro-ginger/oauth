package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/account/profile/domain/profile"
)

func (repo *repo[T]) List(q query.Query) ([]*profile.Profile[T], errors.Error) {
	q = query.NewModelsQuery(q).
		WithModelsHandlerFunc(func() any {
			return new([]*profile.Profile[T])
		})
	r, err := repo.Repository.List(q)
	if err != nil {
		return nil, err
	}
	return *r.(*[]*profile.Profile[T]), nil
}

func (repo *repo[T]) Get(q query.Query) (*profile.Profile[T], errors.Error) {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(profile.Profile[T])
		})
	r, err := repo.Repository.Get(q)
	if err != nil {
		return nil, err
	}
	return r.(*profile.Profile[T]), nil
}
