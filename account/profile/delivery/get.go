package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/log/logger"
	"github.com/micro-blonde/auth/profile"
	"github.com/micro-blonde/file"
	fileClient "github.com/micro-blonde/file/client"
	pd "github.com/micro-ginger/oauth/account/profile/domain/delivery/profile"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type GetHandler[T profile.Model, F file.Model] interface {
	gateway.Handler
	Initialize(file fileClient.Client[F])
}

type get[T profile.Model, F file.Model] struct {
	gateway.Responder
	logger log.Logger
	uc     p.UseCase[T]
	file   fileClient.Client[F]
}

func NewGet[T profile.Model, F file.Model](logger log.Logger,
	uc p.UseCase[T], responder gateway.Responder) GetHandler[T, F] {
	h := &get[T, F]{
		Responder: responder,
		logger:    logger,
		uc:        uc,
	}
	return h
}

func (h *get[T, F]) Initialize(file fileClient.Client[F]) {
	h.file = file
}

func (h *get[T, F]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()

	accId := request.GetAuthorization().GetApplicantId().(uint64)
	prof, err := h.uc.GetById(ctx, accId)
	if err != nil {
		if err.IsType(errors.TypeNotFound) {
			prof = new(p.Profile[T])
			return pd.NewProfile(prof), nil
		}
		return nil, err
	}
	if prof.Photo != "" {
		url, err := h.file.GetDownloadUrlByKey(prof.Photo)
		if err != nil {
			h.logger.
				With(logger.Field{
					"error": err.Error(),
				}).
				WithTrace("file.GetDownloadUrlByKey").
				Errorf("error on get download url by key")
			prof.Photo = ""
		} else {
			prof.Photo = url
		}
	}

	return pd.NewProfile(prof), nil
}
