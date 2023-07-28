package handler

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *handler[acc]) Set(ctx context.Context,
	challenge, key string, value any) errors.Error {
	if err := h.cache.SetItem(ctx,
		h.getChallengeKey(challenge), key, value); err != nil {
		return err
	}
	return nil
}

func (h *handler[acc]) Save(ctx context.Context,
	sess *session.Session[acc]) errors.Error {
	challenge, err := h.challengeGenerator(
		h.config.Challenge.Characters,
		h.config.Challenge.Length,
	)
	if err != nil {
		return err
	}

	if sess.Challenge != "" {
		err := h.cache.Rename(ctx,
			h.getChallengeKey(sess.Challenge),
			h.getChallengeKey(challenge))
		if err != nil {
			return err.
				WithTrace("cache.Rename")
		}
	}

	sess.Challenge = challenge

	err = h.cache.Store(ctx, challenge, sess, h.config.Expiration)
	if err != nil {
		return err.WithTrace("Set")
	}
	return nil
}
