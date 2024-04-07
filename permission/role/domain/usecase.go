package domain

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/role/domain/role"
)

type UseCase interface {
	Create(ctx context.Context, item *role.Role) errors.Error
	Count(ctx context.Context, q query.Query) (uint64, errors.Error)
	List(ctx context.Context, q query.Query) ([]*role.Role, errors.Error)
	ListByNames(ctx context.Context, names []string) ([]*role.Role, errors.Error)
	GetById(ctx context.Context, id uint64) (*role.Role, errors.Error)
	GetByName(ctx context.Context, name string) (*role.Role, errors.Error)
	UpdateById(ctx context.Context,
		id uint64, update *role.UpdateRequest) errors.Error
	DeleteById(ctx context.Context, id uint64) errors.Error
}
