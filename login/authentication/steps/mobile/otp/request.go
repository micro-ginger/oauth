package otp

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/response"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

type body struct {
	Mobile *string `json:"mobile"`
}

func (h *_handler[acc]) request(ctx context.Context, request gateway.Request,
	sess *session.Session[acc]) (response.Response, errors.Error) {
	body := new(body)
	if err := request.ProcessBody(body); err != nil {
		return nil, err.
			WithTrace("request.ProcessBody")
	}
	if body.Mobile != nil && len(*body.Mobile) > 0 {
		sess.Info.SetTemp("mobile", *body.Mobile)
	}

	if err := h.Base.Session.Save(ctx, sess); err != nil {
		return nil, err.
			WithTrace("Base.Session.Save")
	}

	return h.Handler.Process(request, sess)
}
