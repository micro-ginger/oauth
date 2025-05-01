package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/session/domain/session"
)

func (uc *useCase[AccountDetail]) Delete(
	ctx context.Context, accId uint64, id string,
) errors.Error {
	return uc.repo.Delete(ctx, uc.getSessionKey(accId, id))
}

func (uc *useCase[AccountDetail]) DeleteAccess(
	ctx context.Context, token string,
) errors.Error {
	return uc.repo.DeleteAccess(ctx, uc.getAccessKey(token))
}

func (uc *useCase[AccountDetail]) DeleteRefresh(
	ctx context.Context, token string,
) errors.Error {
	return uc.repo.DeleteRefresh(ctx, uc.getRefreshKey(token))
}

func (uc *useCase[AccountDetail]) DeleteAll(
	ctx context.Context, session *session.Session[AccountDetail],
) errors.Error {
	if err := uc.DeleteRefresh(ctx, session.RefreshToken); err != nil {
		return err
	}
	if err := uc.DeleteAccess(ctx, session.AccessToken); err != nil {
		return err
	}
	if err := uc.Delete(ctx, session.Account.Id, session.Id); err != nil {
		return err
	}
	return nil
}
