package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

func (uc *useCase) Count(ctx context.Context,
	q query.Query) (uint64, errors.Error) {
	count, err := uc.repo.Count(q)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (uc *useCase) Paginate(ctx context.Context,
	q query.Query) ([]*scope.Scope, errors.Error) {
	items, err := uc.repo.ReadItems(q)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (uc *useCase) Get(ctx context.Context,
	id uint64) (*scope.Scope, errors.Error) {
	q := query.NewFilter(query.New(ctx)).
		WithId(id)
	item, err := uc.repo.ReadItem(q)
	if err != nil {
		return nil, err
	}
	return item, nil
}
