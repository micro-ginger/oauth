package grpc

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/errors/grpc"
	"github.com/ginger-core/log"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/profile"
	prof "github.com/micro-blonde/auth/proto/auth/account/profile"
	profDlv "github.com/micro-ginger/oauth/account/profile/domain/delivery/profile"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type list[T profile.Model] struct {
	logger log.Logger
	uc     p.UseCase[T]
}

func NewList[T profile.Model](logger log.Logger,
	uc p.UseCase[T]) p.GrpcProfilesGetter {
	h := &list[T]{
		logger: logger,
		uc:     uc,
	}
	return h
}

func (h *list[T]) ListProfiles(ctx context.Context,
	request *prof.ListRequest) (*prof.Profiles, error) {
	r, err := h.listProfiles(ctx, request)
	if err != nil {
		h.logger.
			WithContext(ctx).
			With(logger.Field{
				"error": err.Error(),
				"ids":   request.Ids,
			}).
			Errorf("accounts request")
		return r, grpc.Generate(err)
	}
	h.logger.
		WithContext(ctx).
		With(logger.Field{
			"ids": request.Ids,
		}).
		Infof("accounts request")
	return r, nil
}

func (h *list[T]) listProfiles(ctx context.Context,
	request *prof.ListRequest) (*prof.Profiles, errors.Error) {
	var err errors.Error
	var ps []*p.Profile[T]
	r := new(prof.Profiles)
	if len(request.Ids) > 0 {
		q := query.NewFilter(query.New(ctx)).
			WithMatch(&query.Match{
				Key:      "id",
				Operator: query.In,
				Value:    request.Ids,
			})
		ps, err = h.uc.List(ctx, q)
		if err != nil {
			return r, err.
				WithTrace("uc.List")
		}
	} else {
		return r, errors.Validation().
			WithMessage("no reference given")
	}
	r, err = profDlv.GetGrpcProfiles(ps)
	if err != nil {
		return nil, err.
			WithTrace("delivery.GetGrpcAccounts")
	}
	return r, nil
}
