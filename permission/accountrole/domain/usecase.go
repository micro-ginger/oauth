package domain

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/permission/accountrole/domain/accountrole"
	"github.com/micro-ginger/oauth/permission/role/domain/role"
)

type UseCase interface {
	RegisterCreateEventHandle(handle accountrole.CreatedRoleEventHandle)

	Create(ctx context.Context, item *accountrole.AccountRole) errors.Error
	Assign(ctx context.Context, accId uint64,
		role string, isAuthorized *bool) errors.Error
	CreateBulk(ctx context.Context,
		accountId uint64, roles accountrole.CreateBulk) errors.Error

	Getaccountroles(ctx context.Context,
		accountId uint64, getAll bool) ([]*role.Detailed, errors.Error)

	Delete(ctx context.Context, accountId uint64, roleId uint64) errors.Error
	DeleteBulk(ctx context.Context, accountId uint64, roleIds []uint64) errors.Error
}