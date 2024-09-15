package otp

import (
	"context"
	"strings"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/response"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

type body struct {
	Mobile *string `json:"mobile"`
}

func (h *_handler[acc]) request(_ context.Context, request gateway.Request,
	sess *session.Session[acc]) (response.Response, errors.Error) {
	body := new(body)
	if err := request.ProcessBody(body); err != nil {
		return nil, err.
			WithTrace("request.ProcessBody")
	}
	if body.Mobile != nil && len(*body.Mobile) > 0 {
		mobile := *body.Mobile
		mobile = strings.ReplaceAll(mobile, " ", "")
		sess.Info.SetTemp("mobile", mobile)
	}

	r, err := h.Handler.Process(request, sess)
	if err != nil {
		return nil, err.
			WithTrace("Handler.Process")
	}
	return r, nil
}
