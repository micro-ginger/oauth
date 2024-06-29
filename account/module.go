package account

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/delivery"
	"github.com/micro-ginger/oauth/account/delivery/grpc"
	a "github.com/micro-ginger/oauth/account/domain/account"
	r "github.com/micro-ginger/oauth/account/repository"
	"github.com/micro-ginger/oauth/account/usecase"
)

type Module[T a.Model] struct {
	Repository a.Repository[T]
	UseCase    a.UseCase[T]

	GetHandler gateway.Handler

	GrpcGetHandler  a.GrpcAccountGetter
	GrpcListHandler a.GrpcAccountsGetter
}

func New[T a.Model](logger log.Logger, registry registry.Registry,
	baseRepo repository.Repository, responder gateway.Responder) *Module[T] {
	repo := r.New[T](baseRepo)
	uc := usecase.New(logger, registry, repo)
	m := &Module[T]{
		Repository: repo,
		UseCase:    uc,
		GetHandler: delivery.NewGet[T](
			logger.WithTrace("delivery.get"), uc, responder,
		),
		GrpcGetHandler:  grpc.NewGet(logger.WithTrace("grpcGet"), uc),
		GrpcListHandler: grpc.NewList(logger.WithTrace("grpcList"), uc),
	}
	return m
}
