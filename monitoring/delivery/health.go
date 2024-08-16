package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
)

type HealthHandler interface {
	gateway.Handler
	Initialize(repositories ...repository.Repository)
}

type health struct {
	gateway.Responder
	logger log.Logger

	repositories []repository.Repository
}

func NewHealth(logger log.Logger, responder gateway.Responder) HealthHandler {
	h := &health{
		Responder: responder,
		logger:    logger,
	}
	return h
}

func (h *health) Initialize(repositories ...repository.Repository) {
	h.repositories = repositories
}

func (h *health) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()
	for _, repo := range h.repositories {
		if err := repo.Ping(ctx); err != nil {
			return nil, err.WithTrace("repo.Ping")
		}
	}
	return nil, nil
}
