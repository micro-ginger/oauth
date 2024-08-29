package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

func (uc *useCase) Create(ctx context.Context,
	item *scope.Scope) errors.Error {
	q := query.New(ctx)
	if err := uc.repo.Create(q, item); err != nil {
		return err.WithTrace("repo.Create")
	}
	return nil
}

func (uc *useCase) CreateBulk(ctx context.Context,
	scopes []*scope.Scope) errors.Error {
	if err := uc.repo.CreateBulk(query.New(ctx), scopes); err != nil {
		return err
	}

	return nil
}
