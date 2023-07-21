package usecase

import (
	"context"

	"github.com/ginger-core/errors"
)

func (uc *useCase) Delete(ctx context.Context, key any) errors.Error {
	return uc.repo.Delete(ctx, uc.getKey(key))
}
