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

type ReadHandler[T profile.Model, F file.Model] interface {
	Initialize(file fileClient.Client[F])
	pd.Read[T, F]
}

type read[T profile.Model, F file.Model] struct {
	logger log.Logger

	file fileClient.Client[F]
}

func NewRead[T profile.Model, F file.Model](logger log.Logger) ReadHandler[T, F] {
	h := &read[T, F]{
		logger: logger,
	}
	return h
}

func (h *read[T, F]) Initialize(file fileClient.Client[F]) {
	h.file = file
}

func (h *read[T, F]) PopulateRead(request gateway.Request,
	prof *p.Profile[T]) errors.Error {
	if prof.Photo != nil {
		url, err := h.file.GetDownloadUrlByKey(*prof.Photo)
		if err != nil {
			h.logger.
				With(logger.Field{
					"error": err.Error(),
				}).
				WithTrace("file.GetDownloadUrlByKey").
				Errorf("error on get download url by key")
			prof.Photo = nil
		} else {
			prof.Photo = &url
		}
	}
	return nil
}
