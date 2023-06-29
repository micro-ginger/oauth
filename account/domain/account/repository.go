package account

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
)

type Repository[T any] interface {
	Create(q query.Query, obj *Account[T]) errors.Error
	Count(q query.Query) (uint64, errors.Error)
	List(q query.Query) ([]*Account[T], errors.Error)
	Get(q query.Query) (*Account[T], errors.Error)
	Update(q query.Query, update *Account[T]) errors.Error
}
