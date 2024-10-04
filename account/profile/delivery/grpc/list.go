package grpc

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/profile"
	prof "github.com/micro-blonde/auth/proto/auth/account/profile"
	"github.com/micro-blonde/file"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type ListHandler[T profile.Model, F file.Model] interface {
	gateway.Handler
}

type list[T profile.Model, F file.Model] struct {
	read BaseReadHandler[T, F]
	uc   p.UseCase[T]
}

func NewList[T profile.Model, F file.Model](logger log.Logger,
	uc p.UseCase[T], read BaseReadHandler[T, F]) ListHandler[T, F] {
	h := &list[T, F]{
		read: read,
		uc:   uc,
	}
	return h
}

func (h *list[T, F]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()
	q := query.New(ctx)
	var err errors.Error
	q, err = request.ProcessFilters(q, h.read.GetInstruction())
	if err != nil {
		return nil, err.WithTrace("request.ProcessFilters")
	}
	profs, err := h.uc.ListAggregated(ctx, q)
	if err != nil {
		return nil, err.
			WithTrace("uc.GetAggregated")
	}
	items := make([]*prof.Profile, len(profs))
	for i, itm := range profs {
		items[i], err = h.read.GetProfile(itm)
		if err != nil {
			return nil, err.WithTrace("getProfile")
		}
	}
	return &prof.Profiles{
		Items: items,
	}, nil
}
