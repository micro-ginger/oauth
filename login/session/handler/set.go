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

	isNew := true
	if sess.Challenge != "" {
		isNew = false
		err := h.cache.Rename(ctx,
			h.getChallengeKey(sess.Challenge),
			h.getChallengeKey(challenge))
		if err != nil {
			return err
		}
	}

	sess.Challenge = challenge

	err = h.Set(ctx, challenge, root, sess)
	if err != nil {
		return err.WithTrace("Set")
	}
	if isNew {
		// set expiration for the first time
		err = h.cache.Expire(ctx,
			h.getChallengeKey(sess.Challenge),
			h.config.Expiration)
		if err != nil {
			return err.WithTrace("cache.Expire")
		}
	}
	return nil
}
