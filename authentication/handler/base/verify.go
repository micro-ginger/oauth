package base

import (
	"context"
	"fmt"

	cctx "github.com/ginger-core/compound/context"
	"github.com/ginger-core/errors"
	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/authentication/handler/handler"
)

func (h *Handler[acc]) CheckVerifyAccount(
	ctx context.Context, a *a.Account[acc]) errors.Error {
	_verfied := ctx.Value(handler.VerifiedKey)
	verfied := _verfied.(*bool)
	if *verfied {
		return nil
	}
	*verfied = true

	v, err := h.SuspendValidator.BeginRequest(ctx, fmt.Sprint(a.Id))
	if err != nil {
		return err
	}
	go h.SuspendValidator.EndRequest(cctx.NewBackgroundContext(ctx), v)

	if a.Status.Is(account.StatusBlocked) {
		return handler.YourAccountBlockedError
	}
	if a.Status.Is(account.StatusDisabled) {
		return handler.YourAccountDisabledError
	}
	return nil
}
