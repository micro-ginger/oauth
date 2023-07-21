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
	rg "github.com/micro-ginger/oauth/register/domain/delivery/register"
	"github.com/micro-ginger/oauth/register/domain/register"
)

type handler[T register.Model, acc account.Model] struct {
	gateway.Responder
	logger log.Logger
	uc     register.UseCase[T, acc]
}

func NewRegister[T register.Model, acc account.Model](logger log.Logger,
	uc register.UseCase[T, acc], responder gateway.Responder) gateway.Handler {
	h := &handler[T, acc]{
		Responder: responder,
		logger:    logger,
		uc:        uc,
	}
	return h
}

func (h *handler[T, acc]) Handle(request gateway.Request) (any, errors.Error) {
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

	req := &register.Request[T, acc]{
		Account: &a.Account[acc]{
			Account: account.Account[acc]{
				Base: account.Base{
					Id: accId,
				},
			},
		},
		Register: &register.Register[T]{
			AccountId: accId,
		},
	}

	err := h.uc.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	return rg.NewRegister(req.Register), nil
}
