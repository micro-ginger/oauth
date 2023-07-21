package usecase

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/log/logger"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

func (uc *useCase) ValidateRequest(ctx context.Context,
	v *validator.Validation) errors.Error {
	now := time.Now().UTC()
	diff := v.RequestedAt.Add(uc.config.RequestRetryLimitDuration).Sub(now)
	if diff > 0 {
		minutes := int(diff.Minutes())
		seconds := int((diff - time.Minute*time.Duration(minutes)).Seconds()) + 1
		err := errors.Validation().
			WithCode(validator.RetryRequestAfterErrorCode).
			WithId("RetryRequestAfter").
			WithMessage("You can retry request after {{.minutes}} minutes and {{.seconds}} seconds").
			WithMessageOne("You can retry request after {{.seconds}} seconds").
			WithProperty("minutes", minutes).
			WithProperty("seconds", seconds).
			WithPluralCount(minutes + 1)
		return err
	}
	for i, entry := range v.Entries {
		if entry.ExpirationTime.Before(now) {
			// validator expired. reset
			e := uc.NewTemplateEntry(i)
			v.Entries[i] = e
		}
		entry.RemainingRequests -= 1
		// validate
		if entry.RemainingRequests < 0 {
			return validator.TooManyRequestsError
		}
	}
	return nil
}

func (uc *useCase) BeginRequest(ctx context.Context, key any) (*validator.Validation, errors.Error) {
	v := uc.repo.Get(ctx, uc.getKey(key))
	if v == nil {
		v = uc.NewTemplate()
	}
	v.Key = key
	if err := uc.ValidateRequest(ctx, v); err != nil {
		return nil, err
	}
	return v, nil
}

func (uc *useCase) Requested(ctx context.Context, v *validator.Validation) {
	v.RequestedAt = time.Now().UTC()
}

func (uc *useCase) EndRequest(ctx context.Context, v *validator.Validation) {
	uc.Requested(ctx, v)
	if err := uc.repo.Store(ctx, uc.getKey(v.Key), v, uc.maxExpirationTime); err != nil {
		uc.logger.
			With(logger.Field{
				"error": err.Error(),
			}).
			WithId("EndRequest.storeErr").
			Warnf("error on store validation")
	}
}
