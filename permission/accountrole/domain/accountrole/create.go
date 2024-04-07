package accountrole

import "context"

type Create struct {
	RoleId       uint64
	IsAuthorized *bool
}

type CreateBulk []Create

type CreatedRoleEventHandle func(ctx context.Context, roleId, accountId uint64)
