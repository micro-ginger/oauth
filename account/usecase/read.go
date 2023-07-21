package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/account/domain/account"
)

func (uc *useCase[T]) Count(ctx context.Context,
	query query.Query) (uint64, errors.Error) {
	return uc.repo.Count(query)
}

func (uc *useCase[T]) List(ctx context.Context,
	query query.Query) ([]*account.Account[T], errors.Error) {
	return uc.repo.List(query)
}

func (uc *useCase[T]) Get(ctx context.Context,
	query query.Query) (*account.Account[T], errors.Error) {
	return uc.repo.Get(query)
}

func (uc *useCase[T]) GetById(ctx context.Context,
	id uint64) (*account.Account[T], errors.Error) {
	q := query.NewFilter(query.New(ctx)).WithId(id)
	return uc.repo.Get(q)
}
