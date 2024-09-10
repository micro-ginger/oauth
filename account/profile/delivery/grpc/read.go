package grpc

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/log"
	"github.com/ginger-core/log/logger"
	"github.com/micro-blonde/auth/profile"
	prof "github.com/micro-blonde/auth/proto/auth/account/profile"
	"github.com/micro-blonde/file"
	fileClient "github.com/micro-blonde/file/client"
	ins "github.com/micro-ginger/oauth/account/profile/delivery/instruction"
	profDlv "github.com/micro-ginger/oauth/account/profile/domain/delivery/profile"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type BaseReadHandler[T profile.Model, F file.Model] interface {
	Initialize(file fileClient.Client[F])
}

type baseRead[T profile.Model, F file.Model] struct {
	instruction instruction.Instruction
	logger      log.Logger
	uc          p.UseCase[T]
	file        fileClient.Client[F]
}

func newBaseRead[T profile.Model, F file.Model](
	logger log.Logger, uc p.UseCase[T]) *baseRead[T, F] {
	h := &baseRead[T, F]{
		instruction: ins.NewInstructionIntegrated(),
		logger:      logger,
		uc:          uc,
	}
	return h
}

func (h *baseRead[T, F]) Initialize(file fileClient.Client[F]) {
	h.file = file
}

func (h *baseRead[T, F]) getProfile(p *p.Profile[T]) (*prof.Profile, errors.Error) {
	r, err := profDlv.GetGrpcProfile[T](p)
	if err != nil {
		return nil, err.
			WithTrace("delivery.GetGrpcAccount")
	}
	if p.Photo != nil {
		url, err := h.file.GetDownloadUrlByKey(*p.Photo)
		if err != nil {
			h.logger.
				With(logger.Field{
					"error": err.Error(),
				}).
				WithTrace("file.GetDownloadUrlByKey").
				Errorf("error on get download url by key")
			p.Photo = nil
		} else {
			p.Photo = &url
		}
	}
	return r, nil
}
