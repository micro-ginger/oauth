package otp

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
)

func (h *handler[acc]) Generate(ctx context.Context,
	key any, challenge string, otpType string) (*Otp, time.Duration, errors.Error) {
	v, err := h.globalValidator.BeginRequest(ctx, key)
	if err != nil {
		return nil, 0, err
	}

	var o = new(Otp)
	err = h.session.GetItem(ctx, challenge, otpType, o)
	if err != nil && !err.IsType(errors.TypeNotFound) {
		return nil, 0, err
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

	if err := h.session.Set(ctx, challenge, otpType, o); err != nil {
		return nil, 0, err
	}

	h.globalValidator.EndRequest(ctx, v)

	h.logger.Debugf("OTP generated. code: %s", o.Code)

	return o, h.config.Validators.Session.RequestExpiration, nil
}
