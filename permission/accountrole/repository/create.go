package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/accountrole/domain/accountrole"
	"gorm.io/gorm"
)

func (repo *repo) Create(q query.Query,
	item *accountrole.AccountRole) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(accountrole.AccountRole)
		})
	err := repo.Repository.Create(q, item)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) CreateBulk(ctx context.Context,
	accountId uint64, roles accountrole.CreateBulk) errors.Error {
	valueStrings := make([]string, len(roles))
	valueArgs := make([]any, len(roles)*4)

	for i, role := range roles {
		valueStrings[i] = "(?, ?, ?)"

		valueArgs[i*2] = accountId
		valueArgs[i*2+1] = role.RoleId
		valueArgs[i*2+2] = role.IsAuthorized
		valueArgs[i*2+3] = role.IsAuthorized
	}

	smt := "INSERT IGNORE INTO account_roles(account_id,role_id,is_authorized) " +
		"VALUES %s ON DUPLICATE KEY UPDATE is_authorized=?"
	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))

	q := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(accountrole.AccountRole))
	if err := q.Exec(smt, valueArgs...).Error; err != nil {
		return errors.New(err).WithTrace("accountrole.CreateBulk.Exec")
	}
	return nil
}

func (repo *repo) Assign(ctx context.Context,
	accId uint64, role string, isAuthorized *bool) errors.Error {
	smt := `INSERT IGNORE INTO account_roles(account_id, role_id, is_authorized)
		SELECT ?, r.id, true
		FROM roles r
		where r.name = ?`

	db := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(accountrole.AccountRole))
	if err := db.Exec(smt, accId, role).Error; err != nil {
		return errors.New(err).WithTrace("db.Exec")
	}
	return nil
}
