package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/role/domain/role"
)

func (uc *useCase) UpdateById(ctx context.Context,
	id uint64, update *role.UpdateRequest) errors.Error {
	query := query.NewFilter(query.New(ctx)).
		WithId(id)
	if err := uc.repo.Update(query, update); err != nil {
		return err
	}
	return nil
}
