package otp

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/log/logger"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

func (h *handler[acc]) Verify(ctx context.Context,
	challenge string, otpType string, code string) (*Otp, errors.Error) {
	o := new(Otp)
	err := h.info.GetItem(ctx, challenge, otpType, o)
	if err != nil {
		return nil, InvalidCodeError.Clone().WithError(err)
	}

	v, err := h.globalValidator.BeginVerify(ctx, o.Key)
	if err != nil {
		if err == validator.RequestExpiredError {
			return nil, CodeExpiredError
		}
		return nil, err
	}
	if err := h.sessionValidator.ValidateVerify(ctx, o.Validation); err != nil {
		if err == validator.RequestExpiredError {
			return nil, CodeExpiredError
		}
		return nil, err
	}

	storeSessionOtp := true
	defer func(key string, otpType string, otp *Otp) {
		h.globalValidator.EndVerify(ctx, v)
		if storeSessionOtp {
			// decrease verification remaining count
			if err := h.info.Set(ctx, key, otpType, otp); err != nil {
				h.logger.
					With(logger.Field{
						"error": err.Error(),
					}).
					WithTrace("Verify.StoreChallengeErr").
					Warnf("error while updating otp remaining verify count")
			}
		}
	}(challenge, otpType, o)
	// check code
	if o.Code != code {
		return nil, InvalidCodeError
	}
	// remove code
	if err := h.info.DeleteItem(ctx, challenge, otpType); err != nil {
		h.logger.
			With(logger.Field{
				"error": err.Error(),
			}).
			WithTrace("Verify.UnsetItemErr").
			Warnf("error while clearing otp code")
		return nil, errors.Internal(err)
	}
	storeSessionOtp = false
	return o, nil
}
