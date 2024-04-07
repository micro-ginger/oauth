package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-repository/sql/query"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

func (uc *useCase) CountScopes(ctx context.Context,
	q query.Query) (uint64, errors.Error) {
	count, err := uc.repo.CountScopes(q)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (uc *useCase) ListScopes(ctx context.Context,
	q query.Query) ([]*scope.Scope, errors.Error) {
	items, err := uc.repo.ListScopes(q)
	if err != nil {
		return nil, err
	}
	return items, nil
}
