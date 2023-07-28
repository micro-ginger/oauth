package verify

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/otp"
	"github.com/micro-ginger/oauth/login/authentication/response"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *_handler[acc]) verify(ctx context.Context, request gateway.Request,
	sess *session.Session[acc]) (response.Response, errors.Error) {
	body := new(verifyBody)
	if err := request.ProcessBody(body); err != nil {
		return nil, err
	}

	_otp, err := h.getOtp(sess)
	if err != nil {
		return nil, err.WithTrace("getOtp")
	}
	if _otp == nil {
		return nil, otp.InvalidCodeError
	}
	err = h.otp.Verify(ctx, _otp, otpType, body.Code)
	if err != nil {
		return nil, err.WithTrace("otp.Verify")
	}

	if body.NationalCode != "" {
		sess.Info.SetTemp("nationalCode", body.NationalCode)
	}

	sess.Flow.Pos.StepIndex++
	return nil, nil
}
