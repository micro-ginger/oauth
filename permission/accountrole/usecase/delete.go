package usecase

import (
	"context"

	"github.com/ginger-core/errors"
)

func (uc *useCase) Delete(ctx context.Context,
	accountId uint64, roleId uint64) errors.Error {
	if err := uc.repo.Delete(ctx, accountId, roleId); err != nil {
		if err.IsType(errors.TypeNotFound) {
			return nil
		}
		return err
	}
	return nil
}

func (uc *useCase) DeleteBulk(ctx context.Context,
	accountId uint64, roleIds []uint64) errors.Error {
	if err := uc.repo.DeleteBulk(ctx, accountId, roleIds); err != nil {
		return err
	}
	return nil
}
