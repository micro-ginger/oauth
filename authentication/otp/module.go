package otp

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/authentication/info"
	"github.com/micro-ginger/oauth/validator"
)

type Module struct {
	Handler Handler

	SessionValidation *validator.Module
	GlobalValidation  *validator.Module
}

func Initialize[acc account.Model](logger log.Logger, registry registry.Registry,
	cache repository.Cache, info info.Handler[acc]) *Module {
	sessionValidator := validator.Initialize(
		logger.WithTrace("validators.session"),
		registry.ValueOf("validators.session"),
		cache,
	)
	globalValidator := validator.Initialize(
		logger.WithTrace("validators.global"),
		registry.ValueOf("validators.global"),
		cache,
	)
	m := &Module{
		Handler:           New(logger, registry, info, sessionValidator.UseCase, globalValidator.UseCase),
		SessionValidation: sessionValidator,
		GlobalValidation:  globalValidator,
	}
	return m
}
