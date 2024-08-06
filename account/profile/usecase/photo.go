package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-blonde/auth/profile"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

func (uc *useCase[T]) UpdatePhoto(ctx context.Context,
	accountId uint64, photoId string) errors.Error {
	profile := &p.Profile[T]{
		Profile: profile.Profile[T]{
			Base: profile.Base{
				Id:    accountId,
				Photo: &photoId,
			},
		},
	}
	if err := uc.Upsert(ctx, profile); err != nil {
		return err.WithTrace("repo.Upsert")
	}
	return nil
}
