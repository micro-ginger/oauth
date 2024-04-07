package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/rolescope/domain/rolescope"
	"gorm.io/gorm"
)

func (repo *repo) Create(q query.Query, item *rolescope.RoleScope) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(rolescope.RoleScope)
		})
	err := repo.Repository.Create(q, item)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) CreateBulk(ctx context.Context,
	roleId uint64, scopeIds []uint64) errors.Error {
	valueStrings := make([]string, len(scopeIds))
	valueArgs := make([]interface{}, len(scopeIds)*2)

	for i, scopeId := range scopeIds {
		valueStrings[i] = "(?, ?)"

		valueArgs[i*2] = roleId
		valueArgs[i*2+1] = scopeId
	}

	smt := "INSERT IGNORE INTO role_scopes(role_id,scope_id) " +
		"VALUES %s ON DUPLICATE KEY UPDATE scope_id=VALUES(scope_id)"
	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))

	q := repo.Repository.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(rolescope.RoleScope))
	if err := q.Exec(smt, valueArgs...).Error; err != nil {
		return errors.New(err)
	}
	return nil
}
