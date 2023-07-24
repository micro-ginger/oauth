package password

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/authentication/info"
	"github.com/micro-ginger/oauth/authentication/steps/base"
	"github.com/micro-ginger/oauth/authentication/steps/handler"
	v "github.com/micro-ginger/oauth/validator"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

type h[acc account.Model] struct {
	*base.Handler[acc]

	logger log.Logger
	config config

	wrongPassValidator validator.UseCase
}

func New[acc account.Model](logger log.Logger, registry registry.Registry,
	base *base.Handler[acc], cache repository.Cache) handler.Handler[acc] {
	h := &h[acc]{
		Handler: base,
		logger:  logger,
	}
	if registry != nil {
		if err := registry.Unmarshal(&h.config); err != nil {
			panic(err)
		}
		wrongPassValidator := v.Initialize(
			logger.WithTrace("validators.wrongPassword"),
			registry.ValueOf("validators.wrongPassword"),
			cache,
		)
		h.wrongPassValidator = wrongPassValidator.UseCase
	}
	h.config.initialize()
	return h
}

func (h *h[acc]) CanStepIn(info *info.Info[acc]) bool {
	return false
}

func (h *h[acc]) CanStepOut(info *info.Info[acc]) bool {
	return true
}

func (h *h[acc]) IsDone(info *info.Info[acc]) bool {
	return true
}
