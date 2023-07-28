package authentication

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication/steps"
	sbase "github.com/micro-ginger/oauth/login/authentication/steps/base"
	keyPw "github.com/micro-ginger/oauth/login/authentication/steps/key/password"
	"github.com/micro-ginger/oauth/login/authentication/steps/password"
	"github.com/micro-ginger/oauth/login/authentication/steps/refresh"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	loginSession "github.com/micro-ginger/oauth/login/session/domain/session"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Base[acc account.Model] struct {
	logger   log.Logger
	registry registry.Registry

	cache repository.Cache

	loginSession loginSession.Handler[acc]

	steps *steps.Module[acc]

	session session.UseCase
}

func NewBase[acc account.Model](logger log.Logger, registry registry.Registry,
	loginSession loginSession.Handler[acc], cache repository.Cache,
	account account.UseCase[acc], session session.UseCase) *Base[acc] {
	m := &Base[acc]{
		logger:       logger,
		registry:     registry,
		loginSession: loginSession,
		cache:        cache,
		session:      session,
	}

	return m
}

func (m *Base[acc]) GetBase() *Base[acc] {
	return m
}

func (m *Base[acc]) Initialize() {
	m.steps = steps.New[acc](m.logger.WithTrace("handlers"))

	m.initializeHandlers()
}

func (m *Base[acc]) initializeHandlers() {
	config := new(config)
	if err := m.registry.Unmarshal(config); err != nil {
		panic(err)
	}
	for key, cfg := range config.Steps {
		m.initializeHandler(
			m.registry.ValueOf("steps."+key), cfg.Type,
		)
	}
}

func (m *Base[acc]) initializeHandler(
	registry registry.Registry, handlerType step.Type) {
	baseHandler := sbase.New(
		m.logger.WithTrace("base"),
		registry.ValueOf("base"),
		m.loginSession,
		m.cache,
	)
	baseHandler.WithType(handlerType)

	var h handler.Handler[acc]
	switch handlerType {
	case keyPw.Type:
		h = keyPw.New(
			m.logger.WithTrace("key_password"),
			registry,
			baseHandler, m.cache,
		)
	case password.Type:
		h = password.New(
			m.logger.WithTrace("password"),
			registry,
			baseHandler,
		)
	case refresh.Type:
		h = refresh.New(
			m.logger.WithTrace("refresh"),
			baseHandler, m.session,
		)
	default:
		m.logger.
			With(logger.Field{
				"type": handlerType,
			}).
			WithTrace("handler.notFound").
			Warnf("step handler not found")
		return
	}
	m.steps.RegisterHandler(handlerType, h)
}

func (m *Base[acc]) GetLoginSession() loginSession.Handler[acc] {
	return m.loginSession
}

func (m *Base[acc]) GetStepHandlers() map[step.Type]handler.Handler[acc] {
	return m.steps.Handlers
}
