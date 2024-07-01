package profile

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/profile"
)

type UseCase[T profile.Model] interface {
	List(ctx context.Context, query query.Query) ([]*Profile[T], errors.Error)
	Get(ctx context.Context, query query.Query) (*Profile[T], errors.Error)
	GetById(ctx context.Context, id uint64) (*Profile[T], errors.Error)
	// GetAggregated returns profile, using account joined query
	GetAggregated(ctx context.Context, query query.Query) (*Profile[T], errors.Error)
	Upsert(ctx context.Context, profile *Profile[T]) errors.Error
	// Load gets or creates the profile and returns the profile as result
	Load(ctx context.Context, id uint64, refId int64) (*Profile[T], errors.Error)
}
