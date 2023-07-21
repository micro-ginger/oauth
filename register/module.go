package register

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/register/delivery"
	"github.com/micro-ginger/oauth/register/domain"
	"github.com/micro-ginger/oauth/register/domain/register"
	r "github.com/micro-ginger/oauth/register/repository"
	"github.com/micro-ginger/oauth/register/usecase"
)

type Module[T register.Model, acc account.Model] struct {
	Repository register.Repository[T]
	UseCase    domain.UseCase[T, acc]

	RegisterHandler gateway.Handler
}

func New[T register.Model, acc account.Model](logger log.Logger,
	baseRepo repository.Transational, responder gateway.Responder) *Module[T, acc] {
	repo := r.New[T](baseRepo)
	uc := usecase.New[T, acc](logger, repo)
	m := &Module[T, acc]{
		Repository: repo,
		UseCase:    uc,
		RegisterHandler: delivery.NewRegister[T, acc](
			logger.WithTrace("delivery.register"), uc, responder,
		),
	}
	return m
}
