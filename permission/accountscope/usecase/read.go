package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
)

func (uc *useCase) GetAllAccountScopes(ctx context.Context,
	accountId uint64, getAll bool) ([]*scope.Detailed, errors.Error) {
	scopes, err := uc.repo.GetAllAccountScopes(ctx, accountId, getAll)
	if err != nil {
		return nil, err
	}
	return scopes, nil
}

func (uc *useCase) GetAccountScopes(ctx context.Context,
	accountId uint64, getAll bool) ([]*scope.Detailed, errors.Error) {
	scopes, err := uc.repo.GetAccountScopes(ctx, accountId, getAll)
	if err != nil {
		return nil, err
	}
	return scopes, nil
}

func (uc *useCase) GetAccountScopesFromRoles(ctx context.Context,
	accountId uint64, roles []string, getAll bool) ([]*scope.Detailed, errors.Error) {
	scopes, err := uc.repo.GetAccountScopesFromRoles(ctx, accountId, roles, getAll)
	if err != nil {
		return nil, err
	}
	return scopes, nil
}

func (uc *useCase) GetAccountScopesFromScopes(ctx context.Context,
	accountId uint64, names []string, getAll bool) ([]*scope.Detailed, errors.Error) {
	scopes, err := uc.repo.GetAccountScopesFromScopes(ctx, accountId, names, getAll)
	if err != nil {
		return nil, err
	}
	return scopes, nil
}

func (uc *useCase) ListDefaultAccountScopes(ctx context.Context,
	accountId uint64, getAll bool) ([]*scope.Detailed, errors.Error) {
	scopes, err := uc.repo.ListDefaultAccountScopes(ctx, accountId, getAll)
	if err != nil {
		return nil, err
	}
	return scopes, nil
}
