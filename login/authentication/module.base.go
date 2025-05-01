package authentication

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication/steps"
	sbase "github.com/micro-ginger/oauth/login/authentication/steps/base"
	keyPw "github.com/micro-ginger/oauth/login/authentication/steps/key/password"
	"github.com/micro-ginger/oauth/login/authentication/steps/password"
	"github.com/micro-ginger/oauth/login/authentication/steps/refresh"
	"github.com/micro-ginger/oauth/login/flow"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	loginSession "github.com/micro-ginger/oauth/login/session/domain/session"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Base[acc account.Model, SessionAccountDetail gateway.ResultGetter] struct {
	logger   log.Logger
	registry registry.Registry

	cache repository.Cache

	account account.UseCase[acc]

	loginSession loginSession.Handler[acc, SessionAccountDetail]

	steps *steps.Module[acc, SessionAccountDetail]

	session session.UseCase[SessionAccountDetail]
}

func NewBase[acc account.Model, SessionAccountDetail gateway.ResultGetter](
	logger log.Logger, registry registry.Registry,
	loginSession loginSession.Handler[acc, SessionAccountDetail],
	cache repository.Cache, account account.UseCase[acc],
	session session.UseCase[SessionAccountDetail],
) *Base[acc, SessionAccountDetail] {
	m := &Base[acc, SessionAccountDetail]{
		logger:       logger,
		registry:     registry,
		loginSession: loginSession,
		cache:        cache,
		account:      account,
		session:      session,
	}

	return m
}

func (m *Base[acc, SessionAccountDetail]) GetBase() *Base[acc, SessionAccountDetail] {
	return m
}

func (m *Base[acc, SessionAccountDetail]) InitializeSteps(flows flow.Flows) {
	m.steps = steps.New[acc, SessionAccountDetail](m.logger.WithTrace("handlers"))
	m.steps.Initialize(flows)
}

func (m *Base[acc, SessionAccountDetail]) Initialize() {
	m.initializeHandlers()
}
func (m *Base[acc, SessionAccountDetail]) initializeHandlers() {
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

func (m *Base[acc, SessionAccountDetail]) initializeHandler(
	registry registry.Registry, handlerType step.Type) {
	baseHandler := sbase.New(
		m.logger.WithTrace("base"),
		registry.ValueOf("base"),
		m.loginSession,
		m.cache,
	)
	baseHandler.
		WithType(handlerType).
		WithAccount(m.account)

	var h handler.Handler[acc, SessionAccountDetail]
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

func (m *Base[acc, SessionAccountDetail]) GetLoginSession() loginSession.Handler[acc, SessionAccountDetail] {
	return m.loginSession
}

func (m *Base[acc, SessionAccountDetail]) GetStepHandlers() map[step.Type]handler.Handler[acc, SessionAccountDetail] {
	return m.steps.Handlers
}
