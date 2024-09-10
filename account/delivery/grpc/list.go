package grpc

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/query"
	acc "github.com/micro-blonde/auth/proto/auth/account"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type ListHandler[T account.Model] interface {
	gateway.Handler
	BaseReadHandler[T]
}

type list[T account.Model] struct {
	*baseRead[T]
}

func NewList[T account.Model](logger log.Logger,
	uc account.UseCase[T]) ListHandler[T] {
	h := &list[T]{
		baseRead: newBaseRead[T](logger, uc),
	}
	return h
}

func (h *list[T]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()
	q := query.New(ctx)
	var err errors.Error
	q, err = request.ProcessFilters(q, h.instruction)
	if err != nil {
		return nil, err.WithTrace("request.ProcessFilters")
	}
	accs, err := h.uc.List(ctx, q)
	if err != nil {
		return nil, err.
			WithTrace("uc.List")
	}
	items := make([]*acc.Account, len(accs))
	for i, acc := range accs {
		items[i], err = h.getAccount(acc)
		if err != nil {
			return nil, err.WithTrace("getAccount")
		}
	}
	return &acc.Accounts{
		Items: items,
	}, nil
}
