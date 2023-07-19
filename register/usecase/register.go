package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/register/domain/register"
)

func (uc *useCase[T, acc]) Register(ctx context.Context,
	request *register.Request[T, acc]) (err errors.Error) {
	q := query.New(ctx)

	err = uc.repo.Begin(q)
	if err != nil {
		return err.
			WithTrace("repo.Begin")
	}
	defer func() {
		if err != nil {
			q.SetError(err)
		}
		err = uc.repo.End(q)
		if err != nil {
			err = err.
				WithTrace("repo.End")
		}
	}()

	err = uc.account.Upsert(ctx, q, request.Account)
	if err != nil {
		return err.WithTrace("account.Upsert")
	}

	request.Register.AccountId = request.Account.Id

	err = uc.repo.Upsert(q, request.Register)
	if err != nil {
		return err.WithTrace("repo.Upsert")
	}

	return nil
}
