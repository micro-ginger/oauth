package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log/logger"
	"github.com/micro-blonde/file"
	fileClient "github.com/micro-blonde/file/client"
	pd "github.com/micro-ginger/oauth/account/profile/domain/delivery/profile"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type PhotoHandler[T file.Model] interface {
	gateway.Handler
	Initialize(file fileClient.Client[T])
}

type photo[T file.Model] struct {
	gateway.Responder
	logger logger.Logger
	uc     p.PhotoUpdater
	file   fileClient.Client[T]
}

func NewUpdatePhoto[T file.Model](logger logger.Logger,
	uc p.PhotoUpdater, responder gateway.Responder) PhotoHandler[T] {
	h := &photo[T]{
		Responder: responder,
		logger:    logger,
		uc:        uc,
	}
	return h
}

func (h *photo[T]) Initialize(file fileClient.Client[T]) {
	h.file = file
}

func (h *photo[T]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()

	ginCtx := request.GetConn().(*gin.Context)

	photo, fErr := ginCtx.FormFile("photo")
	if fErr != nil {
		if fErr == http.ErrMissingFile {
			return nil, errors.Validation(fErr).
				WithId("FileRequiredError").
				WithMessage("file is missing to process your request")
		}
		return nil, errors.Validation(fErr)
	}
	uploadedFile, oErr := photo.Open()
	if oErr != nil {
		return nil, errors.Validation(oErr).
			WithTrace("file.Open")
	}
	data := make([]byte, photo.Size)
	_, rErr := uploadedFile.Read(data)
	if rErr != nil {
		return nil, errors.New(rErr).
			WithTrace("uploadedFile.Read")
	}
	resp, err := h.file.Store(ctx, &file.StoreRequest{
		Data:     data,
		Category: "profile_photo",
		Name:     photo.Filename,
	})
	if err != nil {
		return nil, err.WithTrace("file.Store")
	}
	accId := request.GetAuthorization().GetApplicantId().(uint64)
	err = h.uc.UpdatePhoto(ctx, accId, resp.Id)
	if err != nil {
		return nil, err.
			WithTrace("UpdatePhoto")
	}
	return pd.NewProfilePhoto(resp.Url), nil
}
