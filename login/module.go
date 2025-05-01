package login

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication"
	"github.com/micro-ginger/oauth/login/delivery"
	ad "github.com/micro-ginger/oauth/login/domain/account"
	"github.com/micro-ginger/oauth/login/flow"
	s "github.com/micro-ginger/oauth/login/session/domain/session"
	sessionHandler "github.com/micro-ginger/oauth/login/session/handler"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Module[acc account.Model, SessionAccountDetail gateway.ResultGetter] struct {
	Logger         log.Logger
	Session        s.Handler[acc, SessionAccountDetail]
	Authentication authentication.Module[acc, SessionAccountDetail]
	Handler        delivery.Handler[acc, SessionAccountDetail]
}

func New[acc account.Model, SessionAccountDetail gateway.ResultGetter](
	logger log.Logger, registry registry.Registry,
	account account.UseCase[acc], session session.UseCase[SessionAccountDetail],
	cache repository.Cache, responder gateway.Responder,
) *Module[acc, SessionAccountDetail] {
	sess := sessionHandler.New[acc, SessionAccountDetail](
		logger.WithTrace("session"),
		registry.ValueOf("session"),
		cache,
	)
	auth := authentication.New(
		logger.WithTrace("authentication"),
		registry,
		sess,
		cache,
		account,
		session,
	)
	m := &Module[acc, SessionAccountDetail]{
		Logger:         logger,
		Session:        sess,
		Authentication: auth,
		Handler: delivery.NewLogin[acc, SessionAccountDetail](
			logger.WithTrace("delivery.login"),
			responder,
		),
	}
	return m
}

func (m *Module[acc, SessionAccountDetail]) Initialize(account ad.UseCase[acc],
	flows flow.Flows, session session.UseCase[SessionAccountDetail]) {
	m.Authentication.GetBase().InitializeSteps(flows)
	m.Authentication.Initialize()
	m.Handler.Initialize(account, m.Session, flows, session)

	stepHandlers := m.Authentication.GetStepHandlers()
	for _, h := range stepHandlers {
		m.Handler.RegisterHandler(h.GetType(), h)
	}
}
