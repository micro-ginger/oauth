package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/session/domain/session"
)

func (uc *useCase) Get(ctx context.Context, accId uint64, id string) (*session.Session, errors.Error) {
	return uc.repo.Get(ctx, uc.getSessionKey(accId, id))
}

func (uc *useCase) GetByAccess(ctx context.Context, token string) (*session.Session, errors.Error) {
	return uc.repo.Get(ctx, uc.getAccessKey(token))
}

func (uc *useCase) GetByRefresh(ctx context.Context, token string) (*session.Session, errors.Error) {
	return uc.repo.Get(ctx, uc.getRefreshKey(token))
}

func (uc *useCase) ListSessionIdsOfAccount(ctx context.Context, accId uint64) ([]string, errors.Error) {
	keys, err := uc.repo.ListKeys(ctx, fmt.Sprintf("*_%d_*", accId))
	if err != nil {
		return nil, err
	}
	for i, k := range keys {
		keys[i] = k[strings.LastIndex(k, "_")+1:]
	}
	return keys, nil
}
