package repository

import (
	"fmt"
	"strings"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
	"gorm.io/gorm"
)

func (repo *repo) Create(q query.Query,
	item *scope.Scope) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(scope.Scope)
		})
	err := repo.Repository.Upsert(q, item)
	if err != nil {
		return err.WithTrace("Repository.Upsert")
	}
	return nil
}

func (repo *repo) CreateBulk(q query.Query,
	scopes []*scope.Scope) errors.Error {
	valueStrings := make([]string, len(scopes))
	valueArgs := make([]interface{}, len(scopes)*3)

	for i, scope := range scopes {
		valueStrings[i] = "(?, ?, ?)"

		valueArgs[i*2] = scope.Name
		valueArgs[i*2+1] = scope.State
		valueArgs[i*2+2] = scope.Description
	}

	smt := "INSERT IGNORE INTO scopes(name,state,description) " +
		"VALUES %s"
	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))

	gq := repo.GetDB(q).(*gorm.DB).
		WithContext(q.GetContext()).
		Model(new(scope.Scope))
	if err := gq.Exec(smt, valueArgs...).Error; err != nil {
		return errors.New(err).
			WithTrace("Exec")
	}
	return nil
}
