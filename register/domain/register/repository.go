package register

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/ginger-core/repository"
)

type Repository[T Model] interface {
	Begin(query query.Query, options ...repository.Options) errors.Error
	End(query query.Query) errors.Error

	Get(query query.Query) (*Register[T], errors.Error)
	Upsert(query query.Query, register *Register[T]) errors.Error
	Update(query query.Query, register *Register[T]) errors.Error
}
