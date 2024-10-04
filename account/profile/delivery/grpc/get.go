package grpc

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/profile"
	"github.com/micro-blonde/file"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type GetHandler[T profile.Model, F file.Model] interface {
	gateway.Handler
}

type get[T profile.Model, F file.Model] struct {
	read BaseReadHandler[T, F]
	uc   p.UseCase[T]
}

func NewGet[T profile.Model, F file.Model](logger log.Logger,
	uc p.UseCase[T], read BaseReadHandler[T, F]) GetHandler[T, F] {
	h := &get[T, F]{
		read: read,
		uc:   uc,
	}
	return h
}

func (h *get[T, F]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()
	q := query.New(ctx)
	var err errors.Error
	q, err = request.ProcessFilters(q, h.read.GetInstruction())
	if err != nil {
		return nil, err.WithTrace("request.ProcessFilters")
	}
	prof, err := h.uc.GetAggregated(ctx, q)
	if err != nil {
		return nil, err.
			WithTrace("uc.GetAggregated")
	}
	return h.read.GetProfile(prof)
}
