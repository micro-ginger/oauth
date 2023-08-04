package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/account/domain/account"
)

func (uc *useCase[T]) Upsert(ctx context.Context,
	q query.Query, update *account.Account[T]) errors.Error {
	if err := uc.repo.Upsert(q, update); err != nil {
		return err.WithTrace("repo.Update")
	}
	return nil
}
