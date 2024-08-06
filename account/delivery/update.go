package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
	ad "github.com/micro-ginger/oauth/account/domain/delivery/account"
)

type update[T account.Model] struct {
	gateway.Responder
	logger log.Logger
	uc     a.UseCase[T]
}

func NewUpdate[T account.Model](logger log.Logger,
	uc a.UseCase[T], responder gateway.Responder) gateway.Handler {
	h := &update[T]{
		Responder: responder,
		logger:    logger,
		uc:        uc,
	}
	return h
}

func (h *update[T]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()

	body := new(ad.Update[T])
	if err := request.ProcessBody(body); err != nil {
		return nil, errors.Validation(err)
	}

	update := body.GetUpdate()
	q := query.NewFilter(query.New(ctx)).
		WithMatch(&query.Match{
			Key:      "id",
			Operator: query.Equal,
			Value:    request.GetAuthorization().GetApplicantId().(uint64),
		})
	err := h.uc.Update(ctx, q, update)
	if err != nil {
		return nil, err.WithTrace("uc.Update")
	}
	return nil, nil
}
