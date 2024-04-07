package repository

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/permission/role/domain/role"
	"gorm.io/gorm"
)

func (repo *repo) getRows(r *gorm.DB) ([]*role.Detailed, errors.Error) {
	rows, err := r.Rows()
	if err != nil {
		return nil, errors.New(err).WithTrace("accountrole.getRoles.Rows")
	}
	roles := make([]*role.Detailed, 0)
	for rows.Next() {
		role := new(role.Detailed)
		if err := r.ScanRows(rows, role); err != nil {
			return nil, errors.New(err).WithTrace("accountrole.getRows.ScanRows")
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (repo *repo) Getaccountroles(ctx context.Context,
	accountId uint64, getAll bool) ([]*role.Detailed, errors.Error) {
	query := `SELECT DISTINCT r.*, ur.is_authorized AS is_authorized
			FROM (
					 SELECT r.*
					 FROM roles r
							  LEFT JOIN account_roles ur on r.id = ur.role_id
					 WHERE ur.account_id = ?
				 ) AS r
					 LEFT JOIN account_roles ur ON r.id = ur.role_id AND ur.account_id = ?`
	if !getAll {
		query += `
			WHERE ur.is_authorized IS NOT false;`
	}
	r := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(role.Role)).
		Raw(query, accountId, accountId)
	roles, err := repo.getRows(r)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
