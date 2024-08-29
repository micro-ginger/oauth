package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
)

func (uc *useCase) Delete(ctx context.Context, name string) errors.Error {
	if err := uc.repo.Delete(query.New(ctx), name); err != nil {
		if err.IsType(errors.TypeNotFound) {
			return nil
		}
		return err.WithTrace("repo.Delete")
	}
	return nil
}

func (uc *useCase) DeleteBulk(ctx context.Context, names []string) errors.Error {
	if err := uc.repo.DeleteBulk(query.New(ctx), names); err != nil {
		return err.WithTrace("repo.DeleteBulk")
	}
	return nil
}
