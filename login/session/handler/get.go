package handler

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *handler[acc]) Get(ctx context.Context,
	challenge string) (*session.Session[acc], errors.Error) {
	sess := new(session.Session[acc])
	if err := h.cache.Load(ctx, challenge, sess); err != nil {
		return nil, err
	}
	sess.AddState(session.StateFromDB)
	sess.Challenge = challenge
	return sess, nil
}
