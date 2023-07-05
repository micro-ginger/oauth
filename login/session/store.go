package session

import (
	"context"

	"github.com/ginger-core/errors"
)

func (uc *useCase) Store(ctx context.Context,
	session *Session) (err errors.Error) {
	if session.challenge != "" {
		if err = uc.Delete(ctx, session); err != nil {
			return err.WithTrace("Delete")
		}
	}
	// generate new key
	session.challenge, err = uc.challengeGenerator()
	if err != nil {
		return err.WithTrace("challengeGenerator")
	}
	// store
	if err = uc.cache.Store(ctx, session.GetKey(),
		session, uc.config.Expiration); err != nil {
		return err.WithTrace("cache.Store")
	}
	return nil
}
