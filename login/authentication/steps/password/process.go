package password

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/response"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

type body struct {
	Password string `json:"password" binding:"required"`
}

func (h *h[acc]) Process(request gateway.Request,
	sess *session.Session[acc]) (response.Response, errors.Error) {
	ctx := request.GetContext()
	body := new(body)
	if err := request.ProcessBody(body); err != nil {
		return nil, err
	}

	a, err := h.GetAccount(ctx, sess.Info, request, nil)
	if err != nil {
		return nil, err
	}
	// validate
	if err := h.CheckVerifyAccount(ctx, a); err != nil {
		return nil, err
	}

	if err := a.MatchPassword(body.Password); err != nil {
		return nil, errors.Unauthorized(err)
	}

	return nil, nil
}
