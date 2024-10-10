package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	sqlQuery "github.com/ginger-repository/sql/query"
	"github.com/micro-blonde/auth/account"
	ad "github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/register/domain/register"
)

func (uc *useCase[T, acc]) Register(ctx context.Context,
	request *register.Request[T, acc]) (err errors.Error) {
	q := query.New(ctx)
	q = sqlQuery.New(q, nil)

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

	q = query.NewFilter(q).
		WithMatch(&query.Match{
			Key:      "id",
			Operator: query.Equal,
			Value:    request.Register.AccountId,
		})
	if request.Update.UpdateStatus == nil {
		request.Update.UpdateStatus = new(ad.UpdateStatus)
	}
	request.Update.UpdateStatus.Add(account.StatusRegistered)

	err = uc.account.Update(ctx, q, request.Update)
	if err != nil {
		return err.WithTrace("account.Update")
	}

	err = uc.repo.Upsert(q, request.Register)
	if err != nil {
		return err.WithTrace("repo.Upsert")
	}

	return nil
}
