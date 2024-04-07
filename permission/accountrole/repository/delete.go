package repository

import (
	"context"
	"fmt"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/accountrole/domain/accountrole"
	"gorm.io/gorm"
)

func (repo *repo) Delete(ctx context.Context,
	accountId uint64, roleId uint64) errors.Error {
	q := query.New(ctx)
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(accountrole.AccountRole)
		})
	q = query.NewFilter(q).
		WithMatch(&query.Match{
			Key:      "account_id",
			Operator: query.Equal,
			Value:    accountId,
		}).
		WithMatch(&query.Match{
			Key:      "role_id",
			Operator: query.Equal,
			Value:    roleId,
		})
	if err := repo.Repository.Delete(q); err != nil {
		return err
	}
	return nil
}

func (repo *repo) DeleteBulk(ctx context.Context,
	roleId uint64, roleIds []uint64) errors.Error {
	arr := ""
	for _, id := range roleIds {
		arr += fmt.Sprint(id) + ","
	}

	smt := `DELETE FROM account_roles WHERE account_id=? and role_id IN (%s)`
	smt = fmt.Sprintf(smt, arr[:len(arr)-1])

	q := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(accountrole.AccountRole))
	if err := q.Exec(smt, roleId).Error; err != nil {
		return errors.New(err).WithTrace("accountrole.DeleteBulk.Exec")
	}
	return nil
}
