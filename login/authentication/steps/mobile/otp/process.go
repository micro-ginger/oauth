package otp

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/info"
	"github.com/micro-ginger/oauth/login/authentication/response"
)

func (h *_handler[acc]) Process(ctx context.Context, request gateway.Request,
	info *info.Info[acc]) (response.Response, errors.Error) {
	if info.StepInd == 0 {
		return h.request(ctx, request, info)
	}
	return h.Handler.Process(ctx, request, info)
}
