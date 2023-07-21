package repository

import (
	"context"

	"github.com/micro-ginger/oauth/validator/domain/validator"
)

func (repo *repo) Get(ctx context.Context, key string) *validator.Validation {
	v := new(validator.Validation)
	err := repo.cache.Load(ctx, key, v)
	if err != nil {
		return nil
	}
	return v
}
