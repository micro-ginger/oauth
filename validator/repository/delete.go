package repository

import (
	"context"

	"github.com/ginger-core/errors"
)

func (repo *repo) Delete(ctx context.Context, key string) errors.Error {
	return repo.cache.Delete(ctx, key)
}
