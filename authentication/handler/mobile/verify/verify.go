package verify

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/authentication/info"
	"github.com/micro-ginger/oauth/authentication/response"
)

func (h *_handler[acc]) verify(ctx context.Context, request gateway.Request,
	inf *info.Info[acc]) (response.Response, errors.Error) {
	body := new(verifyBody)
	if err := request.ProcessBody(body); err != nil {
		return nil, err
	}

	_, err := h.otp.Verify(ctx, inf.Challenge, otpType, body.Code)
	if err != nil {
		return nil, err
	}

	if body.NationalCode != "" {
		inf.SetTemp("nationalCode", body.NationalCode)
	}

	inf.StepInd++
	return nil, nil
}
