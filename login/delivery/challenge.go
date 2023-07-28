package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *lh[acc]) challenge(request gateway.Request,
	challenge string) (*session.Session[acc], any, errors.Error) {
	sess, err := h.loginSession.Get(request.GetContext(), challenge)
	if err != nil {
		return nil, nil, errors.Unauthorized(err).
			WithTrace("loginSession.Get")
	}

	r, err := h.process(request, sess)
	if err != nil {
		return nil, nil, err.
			WithTrace("process")
	}
	return sess, r, nil
}
