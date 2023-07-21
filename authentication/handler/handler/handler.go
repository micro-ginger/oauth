package handler

import (
	"context"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/authentication/info"
	"github.com/micro-ginger/oauth/authentication/response"
)

type Handler[acc account.Model] interface {
	Clone() Handler[acc]
	WithConfig(registry registry.Registry) Handler[acc]
	WithAccount(account account.UseCase[acc]) Handler[acc]
	WithAccountGetter(getter AccountGetter[acc]) Handler[acc]

	Process(ctx context.Context, request gateway.Request,
		info *info.Info[acc]) (response.Response, errors.Error)

	CanStepIn(info *info.Info[acc]) bool
	CanStepOut(info *info.Info[acc]) bool
	IsDone(info *info.Info[acc]) bool
}
