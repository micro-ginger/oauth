package verify

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/response"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *_handler[acc]) verify(ctx context.Context, request gateway.Request,
	sess *session.Session[acc]) (response.Response, errors.Error) {
	body := new(verifyBody)
	if err := request.ProcessBody(body); err != nil {
		return nil, err
	}

	_, err := h.otp.Verify(ctx, sess.Challenge, otpType, body.Code)
	if err != nil {
		return nil, err
	}

	if body.NationalCode != "" {
		sess.Info.SetTemp("nationalCode", body.NationalCode)
	}

	sess.Flow.Pos.StepIndex++
	return nil, nil
}
