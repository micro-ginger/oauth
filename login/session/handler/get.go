package handler

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *handler[acc, SessionAccountDetail]) Get(ctx context.Context,
	challenge string) (*session.Session[acc, SessionAccountDetail], errors.Error) {
	sess := new(session.Session[acc, SessionAccountDetail])
	if err := h.cache.Load(ctx,
		h.getChallengeKey(challenge), sess); err != nil {
		return nil, err
	}
	sess.AddState(session.StateFromDB)
	sess.Challenge = challenge
	return sess, nil
}
