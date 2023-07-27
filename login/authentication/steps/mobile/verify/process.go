package verify

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/info"
	"github.com/micro-ginger/oauth/login/authentication/response"
)

type verifyBody struct {
	Code         string `json:"code" binding:"required"`
	NationalCode string `json:"nationalCode"`
}

func (h *_handler[acc]) Process(ctx context.Context, request gateway.Request,
	info *info.Info[acc]) (response.Response, errors.Error) {
	if info.StepInd == 0 {
		return h.request(ctx, request, info)
	}
	r, err := h.verify(ctx, request, info)
	if err != nil {
		return nil, err
	}
	return r, nil
}
