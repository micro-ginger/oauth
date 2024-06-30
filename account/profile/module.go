package profile

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-blonde/auth/profile"
	"github.com/micro-ginger/oauth/account/profile/delivery"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
	r "github.com/micro-ginger/oauth/account/profile/repository"
	"github.com/micro-ginger/oauth/account/profile/usecase"
)

type Module[T profile.Model] struct {
	Repository p.Repository[T]
	UseCase    p.UseCase[T]

	GetHandler gateway.Handler
}

func New[T profile.Model](logger log.Logger,
	baseRepo repository.Repository, responder gateway.Responder) *Module[T] {
	repo := r.New[T](baseRepo)
	uc := usecase.New(logger, repo)
	m := &Module[T]{
		Repository: repo,
		UseCase:    uc,
		GetHandler: delivery.NewGet(
			logger.WithTrace("delivery.get"), uc, responder,
		),
	}
	return m
}
