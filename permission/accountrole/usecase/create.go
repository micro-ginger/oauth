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
			break
		}
		return err
	}
	return nil
}

func (uc *useCase) CreateBulk(ctx context.Context,
	accountId uint64, roles accountrole.CreateBulk) errors.Error {
	if err := uc.repo.CreateBulk(ctx, accountId, roles); err != nil {
		return err
	}

	for _, r := range roles {
		if r.IsAuthorized != nil && !*r.IsAuthorized {
			// an item unauthorized, trigger events
			for _, h := range uc.createRoleHandlers {
				h(ctx, r.RoleId, accountId)
			}
			break
		}
	}
	return nil
}

func (uc *useCase) Assign(ctx context.Context,
	accId uint64, role string, isAuthorized *bool) errors.Error {
	if err := uc.repo.Assign(ctx, accId, role, isAuthorized); err != nil {
		return err.WithTrace("repo.Assign")
	}
	return nil
}
