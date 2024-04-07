package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/go-sql-driver/mysql"
	"github.com/micro-ginger/oauth/permission/role/domain/role"
)

func (uc *useCase) Create(ctx context.Context, item *role.Role) errors.Error {
	query := query.New(ctx)
	err := uc.repo.Create(query, item)
	if err != nil {
		switch err := err.GetError().(type) {
		case *mysql.MySQLError:
			if err.Number == 1062 {
				// duplicate
				return errors.Validation(err).
					WithId("DuplicateError").
					WithMessage("This item already exists")
			}
		}
		return err
	}
	return nil
}
