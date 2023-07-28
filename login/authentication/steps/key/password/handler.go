package password

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication/steps/base"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	"github.com/micro-ginger/oauth/login/session/domain/session"
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
		wrongPassValidator := v.New(
			logger.WithTrace("validators.wrongPassword"),
			registry.ValueOf("validators.wrongPassword"),
			cache,
		)
		h.wrongPassValidator = wrongPassValidator.UseCase
	}
	h.config.initialize()
	return h
}

func (h *h[acc]) CanStepIn(sess *session.Session[acc]) bool {
	return false
}

func (h *h[acc]) CanStepOut(sess *session.Session[acc]) bool {
	return true
}

func (h *h[acc]) IsDone(sess *session.Session[acc]) bool {
	return true
}
