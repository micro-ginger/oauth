package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/go-sql-driver/mysql"
	"github.com/micro-ginger/oauth/permission/accountscope/domain/accountscope"
)

func (uc *useCase) Create(ctx context.Context,
	item *accountscope.AccountScope) errors.Error {
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
	accountId uint64, scopes accountscope.CreateScopeBulk) errors.Error {
	if err := uc.repo.CreateBulk(ctx, accountId, scopes); err != nil {
		return err
	}

	for _, s := range scopes {
		if s.IsAuthorized != nil && !*s.IsAuthorized {
			// an item unauthorized, update existing sessions
			for _, h := range uc.refreshScopeHandlers {
				h(ctx, s.ScopeId, accountId)
			}
			break
		}
	}
	return nil
}
