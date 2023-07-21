package verify

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/authentication/info"
	"github.com/micro-ginger/oauth/authentication/response"
)

type body struct {
	Password string `json:"password" binding:"required"`
}

func (h *h[acc]) Process(ctx context.Context, request gateway.Request,
	inf *info.Info[acc]) (response.Response, errors.Error) {
	body := new(body)
	if err := request.ProcessBody(body); err != nil {
		return nil, err
	}

	a, err := h.GetAccount(ctx, inf, request, nil)
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
