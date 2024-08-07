package profile

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-blonde/auth/profile"
	"github.com/micro-blonde/file"
	fileClient "github.com/micro-blonde/file/client"
	"github.com/micro-ginger/oauth/account/profile/delivery"
	"github.com/micro-ginger/oauth/account/profile/delivery/grpc"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
	r "github.com/micro-ginger/oauth/account/profile/repository"
	"github.com/micro-ginger/oauth/account/profile/usecase"
)

type Module[Prof profile.Model, File file.Model] struct {
	Repository p.Repository[Prof]
	UseCase    p.UseCase[Prof]

	GetHandler    delivery.GetHandler[Prof, File]
	UpdateHandler gateway.Handler

	PhotoUpdateHandler delivery.PhotoHandler[File]

	GrpcListHandler grpc.ListHandler[Prof, File]
	GrpcGetHandler  grpc.GetHandler[Prof, File]
}

func New[Prof profile.Model, File file.Model](logger log.Logger,
	baseRepo repository.Repository, responder gateway.Responder) *Module[Prof, File] {
	repo := r.New[Prof](baseRepo)
	uc := usecase.New(logger, repo)
	m := &Module[Prof, File]{
		Repository: repo,
		UseCase:    uc,
		GetHandler: delivery.NewGet[Prof, File](
			logger.WithTrace("delivery.get"), uc, responder,
		),
		UpdateHandler: delivery.NewUpdate[Prof](
			logger.WithTrace("delivery.update"), uc, responder,
		),
		PhotoUpdateHandler: delivery.NewUpdatePhoto[File](
			logger.WithTrace("delivery.photoUpdate"), uc, responder,
		),
		GrpcListHandler: grpc.NewList[Prof, File](logger.WithTrace("grpcList"), uc),
		GrpcGetHandler:  grpc.NewGet[Prof, File](logger.WithTrace("grpcGet"), uc),
	}
	return m
}

func (m *Module[Prof, File]) Initialize(file fileClient.Client[File]) {
	m.GetHandler.Initialize(file)
	m.PhotoUpdateHandler.Initialize(file)
	m.GrpcGetHandler.Initialize(file)
	m.GrpcListHandler.Initialize(file)
}
