package usecase

import (
	"fmt"
	"time"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

type useCase struct {
	logger log.Logger
	config validator.Config

	repo validator.Repository

	validationTemplate validator.Validation
	maxExpirationTime  time.Duration
}

func New(logger log.Logger, registry registry.Registry,
	repo validator.Repository) validator.UseCase {
	uc := &useCase{
		logger: logger,
		repo:   repo,
	}
	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.Initialize()

	uc.validationTemplate.RemainingVerifies = 99999
	for _, v := range uc.config.Validators {
		ent := &validator.Entry{
			MaxRequests:       v.MaximumRequestsCount,
			MaxVerifies:       v.MaximumVerifiesCount,
			RemainingRequests: v.MaximumRequestsCount,
			RemainingVerifies: v.MaximumVerifiesCount,
		}
		uc.validationTemplate.Entries = append(uc.validationTemplate.Entries, ent)
		if uc.maxExpirationTime < v.HistoryCycleDuration {
			uc.maxExpirationTime = v.HistoryCycleDuration
		}
		if uc.validationTemplate.RemainingVerifies > v.MaximumVerifiesCount {
			uc.validationTemplate.RemainingVerifies = v.MaximumVerifiesCount
		}
	}

	return uc
}

func (uc *useCase) NewTemplateEntry(i int) *validator.Entry {
	e := uc.validationTemplate.Entries[i]
	e.ExpirationTime = time.Now().UTC().Add(uc.config.Validators[i].HistoryCycleDuration)
	return e
}

func (uc *useCase) NewTemplate() *validator.Validation {
	v := uc.validationTemplate.Clone()
	for i, e := range v.Entries {
		e.ExpirationTime = time.Now().UTC().
			Add(uc.config.Validators[i].HistoryCycleDuration)
	}
	return v
}

func (uc *useCase) getKey(key any) string {
	return uc.config.KeyPrefix + fmt.Sprint(key)
}
