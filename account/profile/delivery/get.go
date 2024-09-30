package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/auth/profile"
	"github.com/micro-blonde/file"
	pd "github.com/micro-ginger/oauth/account/profile/domain/delivery/profile"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type GetHandler[T profile.Model, F file.Model] interface {
	gateway.Handler
}

type get[T profile.Model, F file.Model] struct {
	gateway.Responder
	pd.Read[T, F]
	logger log.Logger
	uc     p.UseCase[T]
}

func NewGet[T profile.Model, F file.Model](logger log.Logger,
	uc p.UseCase[T], read pd.Read[T, F],
	responder gateway.Responder) GetHandler[T, F] {
	h := &get[T, F]{
		Responder: responder,
		Read:      read,
		logger:    logger,
		uc:        uc,
	}
	return h
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
	h.PopulateRead(request, prof)
	return pd.NewProfile(prof), nil
}
