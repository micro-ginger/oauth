package otp

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/response"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *_handler[acc]) Process(request gateway.Request,
	sess *session.Session[acc]) (response.Response, errors.Error) {
	ctx := request.GetContext()
	if sess.Flow.Pos.ActionIndex == 0 {
		return h.request(ctx, request, sess)
	}
	return h.Handler.Process(request, sess)
}
