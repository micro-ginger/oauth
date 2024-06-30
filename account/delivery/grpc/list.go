package grpc

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/errors/grpc"
	"github.com/ginger-core/log"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-core/query"
	acc "github.com/micro-blonde/auth/proto/auth/account"
	"github.com/micro-ginger/oauth/account/domain/account"
	accDlv "github.com/micro-ginger/oauth/account/domain/delivery/account"
)

type list[T account.Model] struct {
	logger log.Logger
	uc     account.UseCase[T]
}

func NewList[T account.Model](logger log.Logger,
	uc account.UseCase[T]) account.GrpcAccountsGetter {
	h := &list[T]{
		logger: logger,
		uc:     uc,
	}
	return h
}

func (h *list[T]) listAccounts(ctx context.Context,
	request *acc.ListRequest) (*acc.Accounts, errors.Error) {
	var err errors.Error
	var a []*account.Account[T]
	r := new(acc.Accounts)
	if len(request.Ids) > 0 {
		q := query.NewFilter(query.New(ctx)).
			WithMatch(&query.Match{
				Key:      "id",
				Operator: query.In,
				Value:    request.Ids,
			})
		a, err = h.uc.List(ctx, q)
		if err != nil {
			return r, err.
				WithTrace("uc.List")
		}
	} else {
		return r, errors.Validation().
			WithMessage("no reference given")
	}
	r, err = accDlv.GetGrpcAccounts(a)
	if err != nil {
		return nil, err.
			WithTrace("delivery.GetGrpcAccounts")
	}
	return r, nil
}

func (h *list[T]) ListAccounts(ctx context.Context,
	request *acc.ListRequest) (*acc.Accounts, error) {
	r, err := h.listAccounts(ctx, request)
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
