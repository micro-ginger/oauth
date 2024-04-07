package usecase

import (
	"context"

	"github.com/ginger-core/errors"
)

func (uc *useCase) Delete(ctx context.Context,
	roleId uint64, scopeId uint64) errors.Error {
	if err := uc.repo.Delete(ctx, roleId, scopeId); err != nil {
		if err.IsType(errors.TypeNotFound) {
			return nil
		}
		return err
	}
	return nil
}

func (uc *useCase) DeleteBulk(ctx context.Context,
	roleId uint64, scopeIds []uint64) errors.Error {
	if err := uc.repo.DeleteBulk(ctx, roleId, scopeIds); err != nil {
		return err
	}
	return nil
}
