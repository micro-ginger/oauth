package delivery

import (
	"strconv"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/account"
	"github.com/micro-blonde/auth/authorization"
	a "github.com/micro-ginger/oauth/account/domain/account"
	ad "github.com/micro-ginger/oauth/account/domain/delivery/account"
	"github.com/micro-ginger/oauth/global"
)

type get[T account.Model] struct {
	gateway.Responder
	logger log.Logger
	uc     a.UseCase[T]
}

func NewGet[T account.Model](logger log.Logger,
	uc a.UseCase[T], responder gateway.Responder) gateway.Handler {
	h := &get[T]{
		Responder: responder,
		logger:    logger,
		uc:        uc,
	}
	return h
}

func (h *get[T]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()

	auth := request.GetAuthorization().(authorization.Authorization[T])
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

	q := query.NewFilter(query.New(ctx)).
		WithId(accId)

	acc, err := h.uc.Get(ctx, q)
	if err != nil {
		return nil, err
	}

	return ad.NewAccount(acc), nil
}
