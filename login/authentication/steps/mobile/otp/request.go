package otp

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/info"
	"github.com/micro-ginger/oauth/login/authentication/response"
)

type body struct {
	Mobile *string `json:"mobile"`
}

func (h *_handler[acc]) request(ctx context.Context, request gateway.Request,
	inf *info.Info[acc]) (response.Response, errors.Error) {
	body := new(body)
	if err := request.ProcessBody(body); err != nil {
		return nil, err
	}
	if body.Mobile != nil && len(*body.Mobile) > 0 {
		inf.SetTemp("mobile", *body.Mobile)
	}

	if err := h.Base.Info.Save(ctx, inf); err != nil {
		return nil, err
	}

	return h.Handler.Process(ctx, request, inf)
}
