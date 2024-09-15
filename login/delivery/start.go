package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-blonde/auth/authorization"
	ldd "github.com/micro-ginger/oauth/login/domain/delivery/login"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *lh[acc]) start(request gateway.Request) (*session.Session[acc], any, errors.Error) {
	req := new(ldd.Request)
	if err := request.ProcessQueries(req); err != nil {
		return nil, nil, errors.
			Validation(err).
			WithTrace("request.ProcessQueries")
	}
	flow, err := h.getFlow(request, req)
	if err != nil {
		return nil, nil, err.WithTrace("getFlow")
	}
	if flow.Stages[0].Steps[0].IsCaptchaRequired {
		auth := request.GetAuthorization().(authorization.Authorization[acc])
		if !auth.IsCaptchaVerified() {
			return nil, nil, errors.Validation().
				WithTrace("!auth.IsCaptchaVerified").
				WithId("InvalidCaptcha").
				WithMessage("an error occured while verifiying Captcha")
		}
	}
	if h.manager != nil {
		if err := h.manager.BeforeStart(request, req, flow); err != nil {
			return nil, nil, err.WithTrace("manager.BeforeStart")
		}
	}
	sess, err := h.generateSession(request, flow, req)
	if err != nil {
		return nil, nil, err.WithTrace("generateSession")
	}

	r, err := h.process(request, sess)
	if err != nil {
		return nil, nil, err.
			WithTrace("process")
	}
	return sess, r, nil
}
