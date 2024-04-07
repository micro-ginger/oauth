package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/permission/role/domain/role"
)

func (uc *useCase) Getaccountroles(ctx context.Context,
	accountId uint64, getAll bool) ([]*role.Detailed, errors.Error) {
	roles, err := uc.repo.Getaccountroles(ctx, accountId, getAll)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
