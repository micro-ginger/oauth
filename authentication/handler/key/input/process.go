package input

import (
	"context"
	"fmt"

	cctx "github.com/ginger-core/compound/context"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/authentication/handler/handler"
	"github.com/micro-ginger/oauth/authentication/info"
	"github.com/micro-ginger/oauth/authentication/response"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

type body struct {
	Key      string `json:"key" form:"key" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (h *h[acc]) Process(ctx context.Context,
	req gateway.Request, inf *info.Info[acc]) (response.Response, errors.Error) {
	body := new(body)
	if err := req.ProcessBody(body); err != nil {
		return nil, err
	}

	a, err := h.GetAccount(ctx, inf, req, nil)
	if err != nil {
		return nil, handler.InvalidCredentialError.
			Clone().WithError(err)
	}
	// validate
	if err := h.CheckVerifyAccount(ctx, a); err != nil {
		return nil, err
	}
	if err := a.MatchPassword(body.Password); err != nil {
		if h.wrongPassValidator != nil {
			if err := h.checkDisableAccount(ctx, a.Id); err != nil {
				return nil, err.WithTrace("key.input.checkDisableAccount")
			}
		}
		return nil, handler.InvalidCredentialError.Clone().WithError(err)
	}
	//
	inf.PopulateAccount(a)
	return nil, nil
}

func (h *h[acc]) checkDisableAccount(ctx context.Context, accId uint64) errors.Error {
	// cache attempt
	key := fmt.Sprint(accId)
	v, err := h.wrongPassValidator.BeginRequest(ctx, key)
	if err != nil {
		if err.GetCode() == validator.TooManyRequestsErrorCode {
			// disable account
			update := &a.Update[acc]{
				UpdateStatus: &a.UpdateStatus{
					AddStatus: account.StatusDisabled,
				},
			}

			q := query.NewFilter(query.New(ctx)).
				WithId(accId)
			if err = h.Account.Update(ctx, q, update); err != nil {
				return err.WithTrace("checkDisableAccount.updateAcc")
			}
			_ = h.wrongPassValidator.Delete(ctx, key)
			return handler.YourAccountDisabledError
		}
		return err.WithTrace("checkDisableAccount.err")
	}
	go h.wrongPassValidator.EndRequest(cctx.NewBackgroundContext(ctx), v)
	return nil
}
