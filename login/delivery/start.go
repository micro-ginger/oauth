package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	ldd "github.com/micro-ginger/oauth/login/domain/delivery/login"
)

func (h *lh) start(request gateway.Request) (any, errors.Error) {
	req := new(ldd.Request)
	if err := request.ProcessQueries(req); err != nil {
		return nil, errors.
			Validation(err).
			WithTrace("request.ProcessQueries")
	}

	session, err := h.newSession(request, req)
	if err != nil {
		return nil, err.WithTrace("newSession")
	}

	return h.process(request, session)
}
