package otp

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

func (h *handler[acc]) Verify(ctx context.Context,
	o *Otp, otpType string, code string) errors.Error {
	v, err := h.globalValidator.BeginVerify(ctx, o.Key)
	if err != nil {
		if err == validator.RequestExpiredError {
			return CodeExpiredError
		}
		return err.WithTrace("globalValidator.BeginVerify")
	}
	if err := h.sessionValidator.ValidateVerify(ctx, o.Validation); err != nil {
		if err == validator.RequestExpiredError {
			return CodeExpiredError
		}
		return err.WithTrace("sessionValidator.ValidateVerify")
	}

	defer h.globalValidator.EndVerify(ctx, v)
	// check code
	if o.Code != code {
		return InvalidCodeError
	}
	return nil
}
