package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/account/profile/domain/profile"
)

func (uc *useCase[T]) Upsert(ctx context.Context,
	profile *profile.Profile[T]) errors.Error {
	q := query.NewUpdate(query.New(ctx))
	var t T
	if profile.T != t {
		updates := profile.T.GetUpdates()
		for k, v := range updates {
			q.WithSet(k, v)
		}
	}
	if profile.Photo != nil {
		q.WithSet("photo", profile.Photo)
	}
	if err := uc.repo.Upsert(q, profile); err != nil {
		return err
	}
	return nil
}
