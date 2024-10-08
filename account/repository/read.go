package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/account/domain/account"
)

func (repo *repo[T]) Count(q query.Query) (uint64, errors.Error) {
	q = query.NewModelsQuery(q).
		WithModelsHandlerFunc(func() any {
			return new([]*account.Account[T])
		})
	return repo.Repository.Count(q)
}

func (repo *repo[T]) List(q query.Query) ([]*account.Account[T], errors.Error) {
	q = query.NewModelsQuery(q).
		WithModelsHandlerFunc(func() any {
			return new([]*account.Account[T])
		})
	r, err := repo.Repository.List(q)
	if err != nil {
		return nil, err
	}
	return *r.(*[]*account.Account[T]), nil
}

func (repo *repo[T]) Get(q query.Query) (*account.Account[T], errors.Error) {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return account.NewAccount[T]()
		})
	r, err := repo.Repository.Get(q)
	if err != nil {
		return nil, err
	}
	return r.(*account.Account[T]), nil
}
