package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/account/domain/account"
)

func (repo *repo[T]) Upsert(q query.Query,
	update *account.Account[T]) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return account.NewAccount[T]()
		})
	if err := repo.Repository.Upsert(q, update); err != nil {
		return err
	}
	return nil
}
