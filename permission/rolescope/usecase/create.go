package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/go-sql-driver/mysql"
	"github.com/micro-ginger/oauth/permission/rolescope/domain/rolescope"
)

func (uc *useCase) Create(ctx context.Context,
	item *rolescope.RoleScope) errors.Error {
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
	roleId uint64, scopeIds []uint64) errors.Error {
	if err := uc.repo.CreateBulk(ctx, roleId, scopeIds); err != nil {
		return err
	}
	return nil
}
