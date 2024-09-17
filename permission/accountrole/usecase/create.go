package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/go-sql-driver/mysql"
	"github.com/micro-ginger/oauth/permission/accountrole/domain/accountrole"
)

func (uc *useCase) Create(ctx context.Context,
	item *accountrole.AccountRole) errors.Error {
	q := query.New(ctx)
	if err := uc.repo.Create(q, item); err != nil {
		switch err := err.GetError().(type) {
		case *mysql.MySQLError:
			if err.Number == 1062 {
				// duplicate
				return nil
			}
		}
		return err
	}
	return nil
}

func (uc *useCase) Assign(ctx context.Context,
	accId uint64, roles []string) errors.Error {
	if err := uc.repo.Assign(ctx, accId, roles); err != nil {
		return err.WithTrace("repo.Assign")
	}
	return nil
}
