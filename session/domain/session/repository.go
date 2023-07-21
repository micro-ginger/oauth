package session

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
)

type Repository interface {
	Create(ctx context.Context, id string,
		session *Session, exp time.Duration) errors.Error

	Get(ctx context.Context, key string) (*Session, errors.Error)
	ListKeys(ctx context.Context, pattern string) ([]string, errors.Error)

	Delete(ctx context.Context, key string) errors.Error
	DeleteAccess(ctx context.Context, key string) errors.Error
	DeleteRefresh(ctx context.Context, key string) errors.Error
}
