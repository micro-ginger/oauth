package monitoring

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/monitoring/delivery"
)

type Module struct {
	Health delivery.HealthHandler
}

func New(logger log.Logger, responder gateway.Responder) *Module {
	return &Module{
		Health: delivery.NewHealth(logger.WithTrace("health"), responder),
	}
}

func (m *Module) Initialize(repositories ...repository.Repository) {
	m.Health.Initialize(repositories...)
}
