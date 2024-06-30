package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/auth/profile"
	pd "github.com/micro-ginger/oauth/account/profile/domain/delivery"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type get[T profile.Model] struct {
	gateway.Responder
	logger log.Logger
	uc     p.UseCase[T]
}

func NewGet[T profile.Model](logger log.Logger,
	uc p.UseCase[T], responder gateway.Responder) gateway.Handler {
	h := &get[T]{
		Responder: responder,
		logger:    logger,
		uc:        uc,
	}
	return h
}

func (h *get[T]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()

	accId := request.GetAuthorization().GetApplicantId().(uint64)
	prof, err := h.uc.Get(ctx, accId)
	if err != nil {
		if err.IsType(errors.TypeNotFound) {
			prof = new(p.Profile[T])
			return pd.NewProfile(prof), nil
		}
		return nil, err
	}

	return pd.NewProfile(prof), nil
}
