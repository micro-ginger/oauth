package repository

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

func (repo *repo) Store(ctx context.Context,
	key string, v *validator.Validation, exp time.Duration) errors.Error {
	if err := repo.cache.MarshalStore(ctx, key, v, exp); err != nil {
		return err
	}
	return nil
}
