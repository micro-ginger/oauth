package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
)

func (uc *useCase) DeleteById(ctx context.Context, id uint64) errors.Error {
	query := query.NewFilter(query.New(ctx)).
		WithId(id)
	if err := uc.repo.Delete(query); err != nil {
		return err
	}
	return nil
}
