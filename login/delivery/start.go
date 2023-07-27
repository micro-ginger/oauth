package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	ldd "github.com/micro-ginger/oauth/login/domain/delivery/login"
	"github.com/micro-ginger/oauth/login/session"
)

func (h *lh[acc]) start(request gateway.Request) (*session.Session, any, errors.Error) {
	req := new(ldd.Request)
	if err := request.ProcessQueries(req); err != nil {
		return nil, nil, errors.
			Validation(err).
			WithTrace("request.ProcessQueries")
	}

	session, err := h.newSession(request, req)
	if err != nil {
		return nil, nil, err.WithTrace("newSession")
	}

	r, err := h.process(request, session)
	return session, r, err
}
