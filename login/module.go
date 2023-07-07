package login

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/login/delivery"
	"github.com/micro-ginger/oauth/login/flow"
	"github.com/micro-ginger/oauth/login/session"
)

type Module struct {
	Logger  log.Logger
	Session session.UseCase
	Handler delivery.Handler
}

func New(logger log.Logger, registry registry.Registry,
	cache repository.Cache, responder gateway.Responder) *Module {
	sess := session.New(
		logger.WithTrace("session"),
		registry.ValueOf("session"),
		cache,
	)
	m := &Module{
		Logger:  logger,
		Session: sess,
		Handler: delivery.NewLogin(
			logger.WithTrace("delivery.get"), responder,
		),
	}
	return m
}

func (m *Module) Initialize(flows flow.Flows) {
	m.Handler.Initialize(flows, m.Session)
}
