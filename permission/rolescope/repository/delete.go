package repository

import (
	"context"
	"fmt"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/rolescope/domain/rolescope"
	"gorm.io/gorm"
)

func (repo *repo) Delete(ctx context.Context,
	roleId uint64, scopeId uint64) errors.Error {
	q := query.New(ctx)
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(rolescope.RoleScope)
		})
	q = query.NewFilter(q).
		WithMatch(&query.Match{
			Key:      "role_id",
			Operator: query.Equal,
			Value:    roleId,
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
	roleId uint64, scopeIds []uint64) errors.Error {
	arr := ""
	for _, id := range scopeIds {
		arr += fmt.Sprint(id) + ","
	}

	smt := `DELETE FROM role_scopes WHERE role_id=? and scope_id IN (%s)`
	smt = fmt.Sprintf(smt, arr[:len(arr)-1])

	q := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(rolescope.RoleScope))
	if err := q.Exec(smt, roleId).Error; err != nil {
		return errors.New(err)
	}
	return nil
}
