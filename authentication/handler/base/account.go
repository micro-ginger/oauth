package base

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/authentication/info"
	"github.com/micro-ginger/oauth/authentication/request"
)

func (h *Handler[acc]) GetAccount(ctx context.Context,
	inf *info.Info[acc], req gateway.Request,
	request request.Request) (*account.Account[acc], errors.Error) {
	if a := ctx.Value("account"); a != nil {
		return a.(*account.Account[acc]), nil
	}

	if h.AccountGetter != nil {
		return h.AccountGetter(ctx, inf)
	}

	if inf != nil && inf.Account != nil {
		return inf.Account, nil
	}

	var a *account.Account[acc]
	var err errors.Error

	if req.IsAuthenticated() {
		auth := req.GetAuthorization()
		a = auth.GetApplicant().(*account.Account[acc])
		a, err = h.Account.GetById(ctx, a.Id)
		if err != nil {
			return nil, err
		}
	}

	if a == nil {
		return nil, errors.NotFound().
			WithTrace("GetAccount.NotFound").
			WithDesc("applicant not found")
	}

	return a, nil
}
