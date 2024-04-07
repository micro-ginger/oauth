package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/accountscope/domain/accountscope"
	"gorm.io/gorm"
)

func (repo *repo) Create(q query.Query,
	item *accountscope.AccountScope) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(accountscope.AccountScope)
		})
	err := repo.Repository.Create(q, item)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) CreateBulk(ctx context.Context,
	accountId uint64, scopes accountscope.CreateScopeBulk) errors.Error {
	valueStrings := make([]string, len(scopes))
	valueArgs := make([]interface{}, len(scopes)*4)

	for i, scope := range scopes {
		valueStrings[i] = "(?, ?, ?)"

		valueArgs[i*2] = accountId
		valueArgs[i*2+1] = scope.ScopeId
		valueArgs[i*2+2] = scope.IsAuthorized
		valueArgs[i*2+3] = scope.IsAuthorized
	}

	smt := "INSERT IGNORE INTO account_scopes(account_id,scope_id,is_authorized) " +
		"VALUES %s ON DUPLICATE KEY UPDATE is_authorized=?"
	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))

	q := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(accountscope.AccountScope))
	if err := q.Exec(smt, valueArgs...).Error; err != nil {
		return errors.New(err).WithTrace("accountscope.CreateBulk.Exec")
	}
	return nil
}
