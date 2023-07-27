package base

import (
	"context"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication/info"
	"github.com/micro-ginger/oauth/login/authentication/response"
	"github.com/micro-ginger/oauth/login/authentication/step"
	v "github.com/micro-ginger/oauth/validator"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

type Handler[acc account.Model] struct {
	logger log.Logger

	Account account.UseCase[acc]
	Info    info.Handler[acc]

	AccountGetter info.AccountGetter[acc]

	SuspendValidator validator.UseCase
}

func New[acc account.Model](logger log.Logger, registry registry.Registry,
	cache repository.Cache, info info.Handler[acc]) *Handler[acc] {
	suspendValidator := v.Initialize(
		logger.WithTrace("validators.suspend"),
		registry.ValueOf("validators.suspend"),
		cache,
	)

	h := &Handler[acc]{
		logger:           logger,
		Info:             info,
		SuspendValidator: suspendValidator.UseCase,
	}
	return h
}

func (h *Handler[acc]) WithAccount(
	account account.UseCase[acc]) step.Handler[acc] {
	h.Account = account
	return h
}

func (h *Handler[acc]) WithAccountGetter(
	getter info.AccountGetter[acc]) step.Handler[acc] {
	h.AccountGetter = getter
	return h
}

func (h *Handler[acc]) Clone() step.Handler[acc] {
	return h
}

func (h *Handler[acc]) WithConfig(registry registry.Registry) step.Handler[acc] {
	return h
}

func (h *Handler[acc]) Process(ctx context.Context, request gateway.Request,
	info *info.Info[acc]) (response.Response, errors.Error) {
	return nil, errors.NotFound()
}

func (h *Handler[acc]) CanStepIn(info *info.Info[acc]) bool {
	return false
}

func (h *Handler[acc]) CanStepOut(info *info.Info[acc]) bool {
	return false
}

func (h *Handler[acc]) IsDone(info *info.Info[acc]) bool {
	return false
}
