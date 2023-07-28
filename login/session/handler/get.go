package handler

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *handler[acc]) GetItem(ctx context.Context,
	challenge, key string, ref any) errors.Error {
	if err := h.cache.GetItem(ctx,
		h.getChallengeKey(challenge), key, ref); err != nil {
		return err
	}
	return nil
}

func (h *handler[acc]) Get(ctx context.Context,
	challenge string) (*session.Session[acc], errors.Error) {
	sess := new(session.Session[acc])
	if err := h.GetItem(ctx, challenge, root, sess); err != nil {
		return nil, err
	}
	sess.Challenge = challenge
	return sess, nil
}
