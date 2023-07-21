package validator

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
)

type Repository interface {
	Store(ctx context.Context, key string, v *Validation, exp time.Duration) errors.Error
	Get(ctx context.Context, key string) *Validation
	Delete(ctx context.Context, key string) errors.Error
}
