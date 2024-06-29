package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/account/profile/domain/profile"
)

func (uc *useCase[T]) List(ctx context.Context,
	query query.Query) ([]*profile.Profile[T], errors.Error) {
	return uc.repo.List(query)
}

func (uc *useCase[T]) Get(ctx context.Context,
	id uint64) (*profile.Profile[T], errors.Error) {
	q := query.NewFilter(query.New(ctx)).WithId(id)
	r, err := uc.repo.Get(q)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (uc *useCase[T]) Load(ctx context.Context,
	id uint64, refId int64) (*profile.Profile[T], errors.Error) {
	q := query.NewFilter(query.New(ctx)).WithId(id)
	r, err := uc.repo.Get(q)
	if err != nil {
		return nil, err.
			WithTrace("repo.Get")
	}
	return r, nil
}
