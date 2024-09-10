package grpc

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type GetHandler[T account.Model] interface {
	gateway.Handler
	BaseReadHandler[T]
}

type get[T account.Model] struct {
	*baseRead[T]
}

func NewGet[T account.Model](logger log.Logger,
	uc account.UseCase[T]) GetHandler[T] {
	h := &get[T]{
		baseRead: newBaseRead[T](logger, uc),
	}
	return h
}

func (h *get[T]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()
	q := query.New(ctx)
	var err errors.Error
	q, err = request.ProcessFilters(q, h.instruction)
	if err != nil {
		return nil, err.WithTrace("request.ProcessFilters")
	}
	acc, err := h.uc.Get(ctx, q)
	if err != nil {
		return nil, err.
			WithTrace("uc.Get")
	}
	return h.getAccount(acc)
}
