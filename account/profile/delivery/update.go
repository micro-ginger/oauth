package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/auth/profile"
	pd "github.com/micro-ginger/oauth/account/profile/domain/delivery/profile"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type update[T profile.Model] struct {
	gateway.Responder
	logger log.Logger
	uc     p.UseCase[T]
}

func NewUpdate[T profile.Model](logger log.Logger,
	uc p.UseCase[T], responder gateway.Responder) gateway.Handler {
	h := &update[T]{
		Responder: responder,
		logger:    logger,
		uc:        uc,
	}
	return h
}

func (h *update[T]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()

	body := new(pd.Update[T])
	if err := request.ProcessBody(body); err != nil {
		return nil, errors.Validation(err)
	}

	prof := body.GetProfile()
	prof.Id = request.GetAuthorization().GetApplicantId().(uint64)
	err := h.uc.Upsert(ctx, prof)
	if err != nil {
		return nil, err.WithTrace("uc.Upsert")
	}
	return nil, nil
}
