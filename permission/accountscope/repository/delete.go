package repository

import (
	"context"
	"fmt"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/accountscope/domain/accountscope"
	"gorm.io/gorm"
)

func (repo *repo) Delete(ctx context.Context,
	accountId uint64, scopeId uint64) errors.Error {
	q := query.New(ctx)
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(accountscope.AccountScope)
		})
	q = query.NewFilter(q).
		WithMatch(&query.Match{
			Key:      "account_id",
			Operator: query.Equal,
			Value:    accountId,
		}).
		WithMatch(&query.Match{
			Key:      "scope_id",
			Operator: query.Equal,
			Value:    scopeId,
		})
	if err := repo.Repository.Delete(q); err != nil {
		return err
	}
	return nil
}

func (repo *repo) DeleteBulk(ctx context.Context,
	accountId uint64, scopeIds []uint64) errors.Error {
	arr := ""
	for _, id := range scopeIds {
		arr += fmt.Sprint(id) + ","
	}

	smt := `DELETE FROM account_scopes WHERE account_id=? and scope_id IN (%s)`
	smt = fmt.Sprintf(smt, arr[:len(arr)-1])

	q := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(accountscope.AccountScope))
	if err := q.Exec(smt, accountId).Error; err != nil {
		return errors.New(err).WithTrace("accountscope.DeleteBulk.Exec")
	}
	return nil
}

func (repo *repo) Revoke(ctx context.Context,
	accountId uint64, scopes ...string) errors.Error {
	smt := "DELETE FROM account_scopes " +
		"WHERE account_id=? AND " +
		"scope_id IN (SELECT id FROM scopes WHERE name IN(?))"
	q := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(accountscope.AccountScope))
	if err := q.Exec(smt, accountId, scopes).Error; err != nil {
		return errors.New(err).WithTrace("accountscope.CreateBulk.Exec")
	}
	return nil
}
