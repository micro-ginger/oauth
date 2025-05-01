package repository

import (
	"context"

	"github.com/ginger-core/errors"
)

func (repo *repo[AccountDetail]) Delete(
	ctx context.Context, key string) errors.Error {
	if err := repo.cache.Delete(ctx, key); err != nil {
		return err.WithTrace("cache.Delete")
	}
	return nil
}

func (repo *repo[AccountDetail]) DeleteAccess(
	ctx context.Context, key string) errors.Error {
	if err := repo.cache.Delete(ctx, key); err != nil {
		return err.WithTrace("cache.Delete")
	}
	return nil
}

func (repo *repo[AccountDetail]) DeleteRefresh(
	ctx context.Context, key string) errors.Error {
	if err := repo.cache.Delete(ctx, key); err != nil {
		return err.WithTrace("cache.Delete")
	}
	return nil
}
