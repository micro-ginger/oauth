package grpc

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/errors/grpc"
	"github.com/ginger-core/log"
	"github.com/ginger-core/log/logger"
	acc "github.com/micro-blonde/auth/proto/auth/account"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/account/domain/delivery"
)

type get[T account.Model] struct {
	logger log.Logger
	uc     account.UseCase[T]
}

func NewGet[T account.Model](logger log.Logger,
	uc account.UseCase[T]) account.GrpcAccountGetter {
	h := &get[T]{
		logger: logger,
		uc:     uc,
	}
	return h
}

func (h *get[T]) getAccount(ctx context.Context,
	request *acc.GetRequest) (*acc.Account, errors.Error) {
	var err errors.Error
	var a *account.Account[T]
	r := new(acc.Account)
	if request.Id > 0 {
		a, err = h.uc.GetById(ctx, request.Id)
		if err != nil {
			return r, err
		}
	} else {
		return r, errors.Validation().
			WithMessage("no reference given")
	}
	r, err = delivery.GetGrpcAccount[T](a)
	if err != nil {
		return nil, err.
			WithTrace("delivery.GetGrpcAccount")
	}
	return r, nil
}

func (h *get[T]) GetAccount(ctx context.Context,
	request *acc.GetRequest) (*acc.Account, error) {
	r, err := h.getAccount(ctx, request)
	if err != nil {
		h.logger.
			WithContext(ctx).
			With(logger.Field{
				"error": err.Error(),
			}).
			Errorf("account request, id: %d", request.Id)
		return r, grpc.Generate(err)
	}
	h.logger.
		WithContext(ctx).
		Infof("account request, id: %d", request.Id)
	return r, nil
}
