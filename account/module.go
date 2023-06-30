package account

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/delivery"
	a "github.com/micro-ginger/oauth/account/domain/account"
	r "github.com/micro-ginger/oauth/account/repository"
	"github.com/micro-ginger/oauth/account/usecase"
)

type Module struct {
	Repository a.Repository[a.Model]
	UseCase    a.UseCase[a.Model]

	GetHandler gateway.Handler
}

func New(logger log.Logger,
	baseRepo repository.Repository, responder gateway.Responder) *Module {
	repo := r.New[a.Model](baseRepo)
	uc := usecase.New(logger, repo)
	m := &Module{
		Repository: repo,
		UseCase:    uc,
		GetHandler: delivery.NewGet[a.Model](
			logger.WithTrace("delivery.get"), uc, responder,
		),
	}
	return m
}
