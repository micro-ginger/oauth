package repository

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
	"gorm.io/gorm"
)

func (repo *repo) getRowsScopes(r *gorm.DB) ([]*scope.Scope, errors.Error) {
	rows, err := r.Rows()
	if err != nil {
		return nil, errors.New(err)
	}
	scopes := make([]*scope.Scope, 0)
	for rows.Next() {
		scope := new(scope.Scope)
		if err := r.ScanRows(rows, scope); err != nil {
			return nil, errors.New(err).
				WithTrace("accountscopes.getRowsScopes.ScanRows")
		}
		scopes = append(scopes, scope)
	}
	return scopes, nil
}

func (repo *repo) getRows(r *gorm.DB) ([]*scope.Detailed, errors.Error) {
	rows, err := r.Rows()
	if err != nil {
		return nil, errors.New(err).WithTrace("accountscopes.getRows")
	}
	scopes := make([]*scope.Detailed, 0)
	for rows.Next() {
		scope := new(scope.Detailed)
		if err := r.ScanRows(rows, scope); err != nil {
			return nil, errors.New(err).WithTrace("accountscope.getRows.ScanRows")
		}
		scopes = append(scopes, scope)
	}
	return scopes, nil
}

// GetAllAccountScopes gets scopes assigned to account in any ways(scope, role or default)
func (repo *repo) GetAllAccountScopes(ctx context.Context,
	accountId uint64, getAll bool) ([]*scope.Detailed, errors.Error) {
	var params []any
	var query string
	query = `SELECT DISTINCT s.*, us.is_authorized AS is_authorized
			FROM (
					 SELECT DISTINCT s.*
					 FROM scopes s
							  LEFT JOIN account_scopes us ON s.id = us.scope_id
			
							  LEFT JOIN role_scopes rs ON s.id = rs.scope_id
							  LEFT JOIN account_roles ur ON rs.role_id = ur.role_id
							  LEFT JOIN roles r ON rs.role_id = r.id
					 WHERE s.state & 1 = 1
						OR r.state & 1 = 1
						OR us.account_id = ?
						OR ur.account_id = ?
				 ) AS s
					 LEFT JOIN account_scopes us ON s.id = us.scope_id AND us.account_id = ?
			
					 LEFT JOIN role_scopes rs ON s.id = rs.scope_id
					 LEFT JOIN account_roles ur ON rs.role_id = ur.role_id AND ur.account_id = ?`
	if !getAll {
		query += `
			WHERE us.is_authorized IS NOT false
			  AND ur.is_authorized IS NOT false;`
	}
	params = []any{accountId, accountId, accountId, accountId}

	r := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(scope.Detailed)).
		Raw(query, params...)
	scopes, err := repo.getRows(r)
	if err != nil {
		return nil, err.WithTrace("GetAllAccountScopes")
	}
	return scopes, nil
}

func (repo *repo) ListUnauthorizedAccountScopes(ctx context.Context,
	accountId uint64) ([]*scope.Scope, errors.Error) {
	var params []any
	query := `SELECT DISTINCT s.*
			FROM (
					 SELECT DISTINCT s.*
					 FROM scopes s
							  LEFT JOIN account_scopes us ON s.id = us.scope_id
			
							  LEFT JOIN role_scopes rs ON s.id = rs.scope_id
							  LEFT JOIN account_roles ur ON rs.role_id = ur.role_id
							  LEFT JOIN roles r ON rs.role_id = r.id
					 WHERE s.state & 1 = 1
						OR r.state & 1 = 1
						OR us.account_id = ?
						OR ur.account_id = ?
				 ) AS s
					 LEFT JOIN account_scopes us ON s.id = us.scope_id AND us.account_id = ?
			
					 LEFT JOIN role_scopes rs ON s.id = rs.scope_id
					 LEFT JOIN account_roles ur ON rs.role_id = ur.role_id AND ur.account_id = ?
			WHERE us.is_authorized IS false
			  OR ur.is_authorized IS false;`
	params = []any{accountId, accountId, accountId, accountId}

	r := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).
		Model(new(scope.Scope)).
		Raw(query, params...)
	scopes, err := repo.getRowsScopes(r)
	if err != nil {
		return nil, err.WithTrace("ListUnauthorizedAccountScopes")
	}
	return scopes, nil
}

// GetAccountScopes gets scopes assigned to account
func (repo *repo) GetAccountScopes(ctx context.Context,
	accountId uint64, getAll bool) ([]*scope.Detailed, errors.Error) {
	var params []any
	var query string
	query = `SELECT s.*, us.is_authorized AS is_authorized
			FROM (
					 SELECT s.*
					 FROM scopes s
							  JOIN account_scopes us on s.id = us.scope_id
					 WHERE us.account_id = ?
				 ) AS s
					 LEFT JOIN account_scopes us ON s.id = us.scope_id AND us.account_id = ?`
	if !getAll {
		query += `
			WHERE us.is_authorized IS NOT false;`
	}
	params = []any{accountId, accountId}

	r := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).
		Model(new(scope.Scope)).
		Raw(query, params...)
	rows, err := r.Rows()
	if err != nil {
		return nil, errors.New(err).WithTrace("GetAccountScopes.Rows")
	}

	scopes := make([]*scope.Detailed, 0)
	for rows.Next() {
		scope := new(scope.Detailed)
		if err := r.ScanRows(rows, scope); err != nil {
			return nil, errors.New(err).WithTrace("GetAccountScopes.ScanRows")
		}
		scopes = append(scopes, scope)
	}
	return scopes, nil
}

func (repo *repo) GetAccountScopesFromRoles(ctx context.Context,
	accountId uint64, roles []string, getAll bool) ([]*scope.Detailed, errors.Error) {
	var query string

	params := []any{accountId, accountId}
	query = `SELECT DISTINCT s.*, us.is_authorized AS is_authorized
			FROM (
					 SELECT s.*, r.name as role_name
					 FROM scopes s
							  LEFT JOIN role_scopes rs on s.id = rs.scope_id
							  LEFT JOIN account_roles ur on rs.role_id = ur.role_id
							  LEFT JOIN account_scopes us on s.id = us.scope_id
							  LEFT JOIN roles r on rs.role_id = r.id
					 WHERE (r.state & 1 = 1 OR ur.account_id = ? OR us.account_id = ?)
					 `
	if len(roles) > 0 {
		query += `AND r.name IN ?`
		params = append(params, roles)
	}
	query += `) AS s
					 LEFT JOIN account_scopes us ON s.id = us.scope_id AND us.account_id = ?
			
					 LEFT JOIN role_scopes rs ON s.id = rs.scope_id
					 LEFT JOIN account_roles ur ON rs.role_id = ur.role_id AND ur.account_id = ?`
	if !getAll {
		query += `
			WHERE us.is_authorized IS NOT false
			  AND ur.is_authorized IS NOT false;`
	}
	params = append(params, accountId, accountId)

	r := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(scope.Scope)).Raw(query, params...)
	rows, err := r.Rows()
	if err != nil {
		return nil, errors.New(err).WithTrace("GetAccountScopesFromRoles.Rows")
	}

	scopes := make([]*scope.Detailed, 0)
	for rows.Next() {
		scope := new(scope.Detailed)
		if err := r.ScanRows(rows, scope); err != nil {
			return nil, errors.New(err).
				WithTrace("GetAccountScopesFromRoles.ScanRows")
		}
		scopes = append(scopes, scope)
	}
	return scopes, nil
}

func (repo *repo) GetAccountScopesFromScopes(ctx context.Context,
	accountId uint64, names []string, getAll bool) ([]*scope.Detailed, errors.Error) {
	var query string
	query = `SELECT DISTINCT s.*, us.is_authorized AS is_authorized
			FROM (
					 SELECT s.*
					 FROM scopes s
							  LEFT JOIN role_scopes rs on s.id = rs.scope_id
							  LEFT JOIN account_roles ur on rs.role_id = ur.role_id
							  LEFT JOIN account_scopes us on s.id = us.scope_id
							  LEFT JOIN roles r on rs.role_id = r.id
					 WHERE (r.state & 1 = 1 OR ur.account_id = ? OR us.account_id = ?)
					   AND s.name IN ?
				 ) AS s
					 LEFT JOIN account_scopes us ON s.id = us.scope_id AND us.account_id = ?
			
					 LEFT JOIN role_scopes rs ON s.id = rs.scope_id
					 LEFT JOIN account_roles ur ON rs.role_id = ur.role_id AND ur.account_id = ?`
	if !getAll {
		query += `
			WHERE us.is_authorized IS NOT false
			  AND ur.is_authorized IS NOT false;`
	}
	params := []any{accountId, accountId, names, accountId, accountId}

	r := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(scope.Scope)).Raw(query, params...)
	rows, err := r.Rows()
	if err != nil {
		return nil, errors.New(err).WithTrace("GetAccountScopesFromScopes.Rows")
	}

	scopes := make([]*scope.Detailed, 0)
	for rows.Next() {
		scope := new(scope.Detailed)
		if err := r.ScanRows(rows, scope); err != nil {
			return nil, errors.New(err).
				WithTrace("GetAccountScopesFromScopes.ScanRows")
		}
		scopes = append(scopes, scope)
	}
	return scopes, nil
}

func (repo *repo) ListDefaultAccountScopes(ctx context.Context,
	accountId uint64, getAll bool) ([]*scope.Detailed, errors.Error) {
	var params []any
	var query string
	query = `SELECT s.*, us.is_authorized AS is_authorized
			FROM (
					 SELECT s.*
					 FROM scopes s
							  LEFT JOIN account_scopes us on s.id = us.scope_id
					 WHERE us.account_id = ? OR s.state & 1 = 1
				 ) AS s
					 LEFT JOIN account_scopes us ON 
					 	s.id = us.scope_id AND us.account_id = 2`
	if !getAll {
		query += `
			WHERE us.is_authorized IS NOT false;`
	}
	params = []any{accountId, accountId}

	r := repo.GetDB(nil).(*gorm.DB).
		WithContext(ctx).Model(new(scope.Scope)).
		Raw(query, params...)
	rows, err := r.Rows()
	if err != nil {
		return nil, errors.New(err).WithTrace("ListDefaultAccountScopes.Rows")
	}

	scopes := make([]*scope.Detailed, 0)
	for rows.Next() {
		scope := new(scope.Detailed)
		if err := r.ScanRows(rows, scope); err != nil {
			return nil, errors.New(err).
				WithTrace("ListDefaultAccountScopes.ScanRows")
		}
		scopes = append(scopes, scope)
	}
	return scopes, nil
}
