package profile

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/profile"
)

type Repository[T profile.Model] interface {
	List(q query.Query) ([]*Profile[T], errors.Error)
	ListAggregated(query query.Query) ([]*Profile[T], errors.Error)
	Get(query query.Query) (*Profile[T], errors.Error)
	GetAggregated(query query.Query) (*Profile[T], errors.Error)
	Upsert(query query.Query, profile *Profile[T]) errors.Error
}
