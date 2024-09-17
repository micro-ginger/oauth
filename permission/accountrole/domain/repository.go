package domain

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/permission/accountrole/domain/accountrole"
	"github.com/micro-ginger/oauth/permission/role/domain/role"
)

type Repository interface {
	Create(q query.Query, item *accountrole.AccountRole) errors.Error
	Assign(ctx context.Context,
		accId uint64, roles []string) errors.Error
	CreateBulk(ctx context.Context,
		accountId uint64, roles accountrole.CreateBulk) errors.Error

	Getaccountroles(ctx context.Context,
		accountId uint64, getAll bool) ([]*role.Detailed, errors.Error)

	Delete(ctx context.Context, accountId uint64, roleId uint64) errors.Error
	DeleteBulk(ctx context.Context, accountId uint64, roleIds []uint64) errors.Error
}
