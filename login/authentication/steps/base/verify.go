package base

import (
	"context"
	"fmt"

	cctx "github.com/ginger-core/compound/context"
	"github.com/ginger-core/errors"
	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication/steps/handler"
)

func (h *Handler[acc]) CheckVerifyKey(
	ctx context.Context, key string) errors.Error {
	_verfied := ctx.Value(handler.VerifiedKey)
	if _verfied != nil {
		if *_verfied.(*bool) {
			return nil
		}
	}

	validation, err := h.SuspendValidator.BeginRequest(ctx, key)
	if err != nil {
		return err.WithTrace("v.BeginRequest")
	}
	go h.SuspendValidator.EndRequest(cctx.NewBackgroundContext(ctx), validation)

	return nil
}

func (h *Handler[acc]) CheckVerifyAccount(
	ctx context.Context, a *a.Account[acc]) errors.Error {
	if err := h.CheckVerifyKey(ctx, fmt.Sprint(a.Id)); err != nil {
		return err.WithTrace("CheckVerifyKey")
	}

	if a.Status.Is(account.StatusBlocked) {
		return handler.YourAccountBlockedError
	}
	if a.Status.Is(account.StatusDisabled) {
		return handler.YourAccountDisabledError
	}

	return nil
}
