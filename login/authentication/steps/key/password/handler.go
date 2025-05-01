package password

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication/steps/base"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	"github.com/micro-ginger/oauth/login/session/domain/session"
	v "github.com/micro-ginger/oauth/validator"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

type h[acc account.Model, SessionAccountDetail gateway.ResultGetter] struct {
	*base.Handler[acc, SessionAccountDetail]

	logger log.Logger
	config config

	wrongPassValidator validator.UseCase
}

func New[acc account.Model, SessionAccountDetail gateway.ResultGetter](
	logger log.Logger, registry registry.Registry,
	base *base.Handler[acc, SessionAccountDetail], cache repository.Cache,
) handler.Handler[acc, SessionAccountDetail] {
	h := &h[acc, SessionAccountDetail]{
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

func (h *h[acc, SessionAccountDetail]) CanStepIn(
	sess *session.Session[acc, SessionAccountDetail]) bool {
	return false
}

func (h *h[acc, SessionAccountDetail]) CanStepOut(
	sess *session.Session[acc, SessionAccountDetail]) bool {
	return true
}

func (h *h[acc, SessionAccountDetail]) IsDone(
	sess *session.Session[acc, SessionAccountDetail]) bool {
	return true
}
