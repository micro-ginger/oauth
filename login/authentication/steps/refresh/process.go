package refresh

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/info"
	"github.com/micro-ginger/oauth/login/authentication/response"
)

type body struct {
	Token string `json:"token" binding:"required"`
}

func (h *h[acc]) Process(ctx context.Context,
	request gateway.Request, info *info.Info[acc]) (response.Response, errors.Error) {
	body := new(body)
	if err := request.ProcessBody(body); err != nil {
		return nil, err
	}
	session, err := h.session.GetByRefresh(ctx, body.Token)
	if err != nil {
		return nil, errors.Unauthorized(err)
	}

	if err = h.session.DeleteAll(ctx, session); err != nil {
		return nil, err
	}

	// FIXME set old session temp
	// info.Session = session

	a, err := h.Account.GetById(ctx, session.Account.Id)
	if err != nil {
		return nil, errors.Unauthorized(err)
	}

	info.PopulateAccount(a)

	return nil, nil
}
