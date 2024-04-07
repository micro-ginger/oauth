package domain

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/role/domain/role"
)

type Repository interface {
	Create(q query.Query, item *role.Role) errors.Error
	Count(q query.Query) (uint64, errors.Error)
	List(q query.Query) ([]*role.Role, errors.Error)
	Get(q query.Query) (*role.Role, errors.Error)
	Update(q query.Query, item *role.UpdateRequest) errors.Error
	Delete(q query.Query) errors.Error
}
