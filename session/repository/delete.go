package repository

import (
	"context"

	"github.com/ginger-core/errors"
)

func (repo *repo) Delete(ctx context.Context, key string) errors.Error {
	if err := repo.cache.Delete(ctx, key); err != nil {
		return err
	}
	return nil
}

func (repo *repo) DeleteAccess(ctx context.Context, key string) errors.Error {
	if err := repo.cache.Delete(ctx, key); err != nil {
		return err
	}
	return nil
}

func (repo *repo) DeleteRefresh(ctx context.Context, key string) errors.Error {
	if err := repo.cache.Delete(ctx, key); err != nil {
		return err
	}
	return nil
}
