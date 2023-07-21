package repository

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/session/domain/session"
)

func (repo *repo) ListKeys(ctx context.Context,
	pattern string) ([]string, errors.Error) {
	return repo.cache.ListKeys(ctx, pattern)
}

func (repo *repo) Get(ctx context.Context,
	key string) (*session.Session, errors.Error) {
	result := new(session.Session)
	if err := repo.cache.Load(ctx, key, result); err != nil {
		return nil, err
	}
	return result, nil
}
