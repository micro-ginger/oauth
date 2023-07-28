package otp

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
)

func (h *handler[acc]) Generate(ctx context.Context,
	key any, o *Otp, otpType string) (*Otp, time.Duration, errors.Error) {
	v, err := h.globalValidator.BeginRequest(ctx, key)
	if err != nil {
		return nil, 0, err
	}

	if o == nil {
		o = new(Otp)
	}
	if o.Key == 0 || o.Key == nil {
		v := h.sessionValidator.NewTemplate()
		v.Key = key

		o = &Otp{
			Key:        key,
			Code:       h.codeGenerator(),
			Validation: v,
		}
	}
	if o.Validation.RemainingVerifies == 0 {
		h.sessionValidator.ResetVerifies(ctx, o.Validation)
		o.Code = h.codeGenerator()
	}

	if err := h.sessionValidator.ValidateRequest(ctx, o.Validation); err != nil {
		return nil, 0, err
	}
	h.sessionValidator.Requested(ctx, o.Validation)

	h.globalValidator.EndRequest(ctx, v)

	h.logger.Debugf("OTP generated. code: %s", o.Code)

	return o, h.config.Validators.Session.RequestExpiration, nil
}
