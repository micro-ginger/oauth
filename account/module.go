package account

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-blonde/auth/profile"
	"github.com/micro-blonde/file"
	fileClient "github.com/micro-blonde/file/client"
	"github.com/micro-ginger/oauth/account/delivery"
	"github.com/micro-ginger/oauth/account/delivery/grpc"
	a "github.com/micro-ginger/oauth/account/domain/account"
	p "github.com/micro-ginger/oauth/account/profile"
	r "github.com/micro-ginger/oauth/account/repository"
	"github.com/micro-ginger/oauth/account/usecase"
)

type Module[Acc a.Model, Prof profile.Model, File file.Model] struct {
	Repository a.Repository[Acc]
	UseCase    a.UseCase[Acc]

	GetHandler    gateway.Handler
	UpdateHandler gateway.Handler

	GrpcGetHandler  grpc.GetHandler[Acc]
	GrpcListHandler grpc.ListHandler[Acc]

	Profile *p.Module[Prof, File]
}

func New[Acc a.Model, Prof profile.Model, File file.Model](logger log.Logger,
	registry registry.Registry, baseRepo repository.Repository,
	responder gateway.Responder) *Module[Acc, Prof, File] {
	repo := r.New[Acc](baseRepo)
	uc := usecase.New(logger, registry, repo)
	m := &Module[Acc, Prof, File]{
		Repository: repo,
		UseCase:    uc,
		GetHandler: delivery.NewGet(
			logger.WithTrace("delivery.get"), uc, responder,
		),
		UpdateHandler: delivery.NewUpdate(
			logger.WithTrace("delivery.update"), uc, responder,
		),
		GrpcGetHandler:  grpc.NewGet(logger.WithTrace("grpcGet"), uc),
		GrpcListHandler: grpc.NewList(logger.WithTrace("grpcList"), uc),
		Profile: p.New[Prof, File](
			logger.WithTrace("profile"),
			baseRepo, responder,
		),
	}
	return m
}

func (m *Module[Acc, Prof, File]) Initialize(file fileClient.Client[File]) {
	m.Profile.Initialize(file)
}
