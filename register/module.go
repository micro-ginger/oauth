package register

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/register/delivery"
	"github.com/micro-ginger/oauth/register/domain"
	ra "github.com/micro-ginger/oauth/register/domain/account"
	rdd "github.com/micro-ginger/oauth/register/domain/delivery/register"
	"github.com/micro-ginger/oauth/register/domain/register"
	r "github.com/micro-ginger/oauth/register/repository"
	"github.com/micro-ginger/oauth/register/usecase"
)

type Module[R rdd.RequestModel, T register.Model, acc account.Model] struct {
	Repository register.Repository[T]
	UseCase    domain.UseCase[T, acc]

	RegisterHandler delivery.Handler[R, T, acc]
}

func New[R rdd.RequestModel, T register.Model, acc account.Model](logger log.Logger,
	baseRepo repository.Transational, responder gateway.Responder) *Module[R, T, acc] {
	repo := r.New[T](baseRepo)
	uc := usecase.New[T, acc](logger, repo)
	m := &Module[R, T, acc]{
		Repository: repo,
		UseCase:    uc,
		RegisterHandler: delivery.NewRegister[R, T, acc](
			logger.WithTrace("delivery.register"), uc, responder,
		),
	}
	return m
}

func (m *Module[R, T, acc]) Initialize(account ra.UseCase[acc]) {
	m.UseCase.Initialize(account)
}
