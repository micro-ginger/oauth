package refresh

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/response"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

type body struct {
	Token string `json:"token" binding:"required"`
}

func (h *h[acc]) Process(request gateway.Request,
	sess *session.Session[acc]) (response.Response, errors.Error) {
	body := new(body)
	if err := request.ProcessBody(body); err != nil {
		return nil, err
	}
	ctx := request.GetContext()
	session, err := h.session.GetByRefresh(ctx, body.Token)
	if err != nil {
		return nil, errors.Unauthorized(err)
	}

	if err = h.session.DeleteAll(ctx, session); err != nil {
		return nil, err
	}

	sess.Info.SetTemp("session", session)

	a, err := h.Account.GetById(ctx, session.Account.Id)
	if err != nil {
		return nil, errors.Unauthorized(err)
	}

	sess.Info.PopulateAccount(a)

	return nil, nil
}
