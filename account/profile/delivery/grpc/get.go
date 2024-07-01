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

type get[T profile.Model] struct {
	logger log.Logger
	uc     p.UseCase[T]
}

func NewGet[T profile.Model](logger log.Logger,
	uc p.UseCase[T]) p.GrpcProfileGetter {
	h := &get[T]{
		logger: logger,
		uc:     uc,
	}
	return h
}

func (h *get[T]) GetProfile(ctx context.Context,
	request *prof.GetRequest) (*prof.Profile, error) {
	r, err := h.getProfile(ctx, request)
	if err != nil {
		h.logger.
			WithContext(ctx).
			With(logger.Field{
				"error": err.Error(),
			}).
			Errorf("profile request, id: %d", request.Id)
		return r, grpc.Generate(err)
	}
	h.logger.
		WithContext(ctx).
		Infof("profile request, id: %d", request.Id)
	return r, nil
}

func (h *get[T]) getProfile(ctx context.Context,
	request *prof.GetRequest) (*prof.Profile, errors.Error) {
	var err errors.Error
	var p *p.Profile[T]
	r := new(prof.Profile)
	if request.Id > 0 {
		p, err = h.uc.GetById(ctx, request.Id)
		if err != nil {
			return r, err.
				WithTrace("uc.GetById")
		}
	} else if request.Key != "" {
		q := query.NewFilter(query.New(ctx)).
			WithMatch(&query.Match{
				Key:      request.Key,
				Operator: query.Equal,
				Value:    request.Val,
			})
		p, err = h.uc.Get(ctx, q)
		if err != nil {
			return r, err.
				WithTrace("uc.Get")
		}
	} else if request.AccountKey != "" {
		q := query.NewFilter(query.New(ctx)).
			WithMatch(&query.Match{
				Key:      request.AccountKey,
				Operator: query.Equal,
				Value:    request.AccountVal,
			})
		p, err = h.uc.GetAggregated(ctx, q)
		if err != nil {
			return r, err.
				WithTrace("uc.Get")
		}
	} else {
		return r, errors.Validation().
			WithMessage("no reference given")
	}
	r, err = profDlv.GetGrpcProfile[T](p)
	if err != nil {
		return nil, err.
			WithTrace("delivery.GetGrpcAccount")
	}
	return r, nil
}
