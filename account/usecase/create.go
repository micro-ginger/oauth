package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/account/domain/account"
)

func (uc *useCase[T]) Create(ctx context.Context,
	account *account.Account[T]) errors.Error {
	q := query.New(ctx)
	if err := uc.repo.Create(q, account); err != nil {
		return err
	}
	return nil
}
