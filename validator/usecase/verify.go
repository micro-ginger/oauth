package usecase

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/log/logger"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

func (uc *useCase) ResetVerifies(ctx context.Context, v *validator.Validation) {
	v.RemainingVerifies = uc.validationTemplate.RemainingVerifies
}

func (uc *useCase) ValidateVerify(ctx context.Context, v *validator.Validation) errors.Error {
	now := time.Now().UTC()
	// check request
	if uc.config.RequestExpiration > 0 {
		maxRequestValidTime := v.RequestedAt.Add(uc.config.RequestExpiration)
		if now.After(maxRequestValidTime) {
			return validator.RequestExpiredError
		}
	}

	for i, entry := range v.Entries {
		if now.After(entry.ExpirationTime) {
			e := *uc.validationTemplate.Entries[i]
			v.Entries[i] = &e
			continue
		}
		entry.RemainingVerifies -= 1
		// validate
		if entry.RemainingVerifies < 0 {
			return validator.TooManyRequestsError
		}
	}
	return nil
}

func (uc *useCase) BeginVerify(ctx context.Context, key any) (*validator.Validation, errors.Error) {
	v := uc.repo.Get(ctx, uc.getKey(key))
	if v == nil {
		return nil, errors.Validation().
			WithDesc("validation not found to validate verify")
	}
	if err := uc.ValidateVerify(ctx, v); err != nil {
		return nil, err
	}
	return v, nil
}

func (uc *useCase) VerifyChecked(ctx context.Context, v *validator.Validation) {
	v.RemainingVerifies -= 1
	v.VerifyRequestedAt = time.Now().UTC()
}

func (uc *useCase) EndVerify(ctx context.Context, v *validator.Validation) {
	uc.VerifyChecked(ctx, v)
	if err := uc.repo.Store(ctx, uc.getKey(v.Key), v, uc.maxExpirationTime); err != nil {
		uc.logger.
			With(logger.Field{
				"error": err.Error(),
			}).
			WithTrace("EndVerify.StoreErr").
			Warnf("error on store validation")
	}
}
