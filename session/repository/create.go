package repository

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/session/domain/session"
)

func (repo *repo) Create(ctx context.Context,
	key string, session *session.Session, exp time.Duration) errors.Error {
	if err := repo.cache.MarshalStore(ctx, key, session, exp); err != nil {
		return err
	}
	return nil
}
