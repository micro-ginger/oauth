package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/role/domain/role"
)

func (uc *useCase) Count(ctx context.Context,
	q query.Query) (uint64, errors.Error) {
	count, err := uc.repo.Count(q)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (uc *useCase) List(ctx context.Context,
	q query.Query) ([]*role.Role, errors.Error) {
	items, err := uc.repo.List(q)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (uc *useCase) ListByNames(ctx context.Context,
	names []string) ([]*role.Role, errors.Error) {
	q := query.New(ctx)
	q = query.NewFilter(q).
		WithMatch(&query.Match{
			Key:      "name",
			Operator: query.In,
			Value:    names,
		})
	q = query.NewPagination(q).
		WithSize(100).
		WithPage(1)
	items, err := uc.repo.List(q)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (uc *useCase) GetById(ctx context.Context,
	id uint64) (*role.Role, errors.Error) {
	q := query.NewFilter(query.New(ctx)).
		WithId(id)
	item, err := uc.repo.Get(q)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (uc *useCase) GetByName(ctx context.Context,
	name string) (*role.Role, errors.Error) {
	q := query.NewFilter(query.New(ctx)).
		WithMatch(&query.Match{
			Key:      "name",
			Operator: query.Equal,
			Value:    name,
		})
	item, err := uc.repo.Get(q)
	if err != nil {
		return nil, err
	}
	return item, nil
}
