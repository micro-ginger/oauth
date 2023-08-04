package delivery

import (
	"strconv"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/auth/account"
	"github.com/micro-blonde/auth/authorization"
	a "github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/global"
	rdd "github.com/micro-ginger/oauth/register/domain/delivery/register"
	"github.com/micro-ginger/oauth/register/domain/register"
)

type Handler[R rdd.RequestModel, T register.Model, acc account.Model] interface {
	gateway.Handler
	SetRequestModelHandler(reqHandler rdd.RequestModelHandler[R, T, acc])
}

type handler[R rdd.RequestModel, T register.Model, acc account.Model] struct {
	gateway.Responder
	logger log.Logger
	uc     register.UseCase[T, acc]

	reqHandler rdd.RequestModelHandler[R, T, acc]
}

func NewRegister[R rdd.RequestModel, T register.Model, acc account.Model](logger log.Logger,
	uc register.UseCase[T, acc], responder gateway.Responder) Handler[R, T, acc] {
	h := &handler[R, T, acc]{
		Responder: responder,
		logger:    logger,
		uc:        uc,
	}
	return h
}

func (h *handler[R, T, acc]) SetRequestModelHandler(
	reqHandler rdd.RequestModelHandler[R, T, acc]) {
	h.reqHandler = reqHandler
}

func (h *handler[R, T, acc]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()

	auth := request.GetAuthorization().(authorization.Authorization[acc])
	accId := auth.GetAccount().Id

	accIdStr := request.GetParam("account_id")
	if accIdStr != "" {
		if !auth.GetSession().HasScope(global.ScopeManageAccounts) {
			return nil, errors.Forbidden().
				WithTrace("!HasScope(global.ScopeManageAccounts)")
		}
		id, _err := strconv.ParseUint(accIdStr, 10, 64)
		if _err != nil {
			return nil, errors.Validation(_err).
				WithTrace("strconv.ParseUint")
		}
		accId = id
	}
	var err errors.Error

	body := h.reqHandler.New()
	if err = request.ProcessBody(body); err != nil {
		return nil, err.
			WithTrace("request.ProcessBody")
	}
	var hashedPassword []byte
	if body.Password != nil {
		hashedPassword, err = a.HashPassword(*body.Password)
		if err != nil {
			return nil, err.
				WithTrace("a.HashPassword")
		}
	}

	req := &register.Request[T, acc]{
		Account: &a.Account[acc]{
			Account: account.Account[acc]{
				Base: account.Base{
					Id: accId,
				},
			},
			HashedPassword: hashedPassword,
		},
		Register: &register.Register[T]{
			AccountId:      accId,
			HashedPassword: hashedPassword,
		},
	}

	if err = h.reqHandler.PopulateRequest(body, req); err != nil {
		return nil, err.WithTrace("reqHandler.PopulateRequest")
	}

	err = h.uc.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	return rdd.NewResponse(req.Register), nil
}
