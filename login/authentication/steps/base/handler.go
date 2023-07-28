package base

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication/response"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	"github.com/micro-ginger/oauth/login/session/domain/info"
	"github.com/micro-ginger/oauth/login/session/domain/session"
	v "github.com/micro-ginger/oauth/validator"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

type Handler[acc account.Model] struct {
	logger log.Logger

	Type step.Type

	Account account.UseCase[acc]
	Session session.Handler[acc]

	AccountGetter info.AccountGetter[acc]

	SuspendValidator validator.UseCase
}

func New[acc account.Model](logger log.Logger, registry registry.Registry,
	session session.Handler[acc], cache repository.Cache) *Handler[acc] {
	// suspendValidator is otp validation which is being
	// validated in each login session
	suspendValidator := v.New(
		logger.WithTrace("validators.suspend"),
		registry.ValueOf("validators.suspend"),
		cache,
	)

	h := &Handler[acc]{
		logger:           logger,
		Session:          session,
		SuspendValidator: suspendValidator.UseCase,
	}
	return h
}

func (h *Handler[acc]) WithType(t step.Type) handler.Handler[acc] {
	h.Type = t
	return h
}

func (h *Handler[acc]) GetType() step.Type {
	return h.Type
}

func (h *Handler[acc]) WithAccount(
	account account.UseCase[acc]) handler.Handler[acc] {
	h.Account = account
	return h
}

func (h *Handler[acc]) WithAccountGetter(
	getter info.AccountGetter[acc]) handler.Handler[acc] {
	h.AccountGetter = getter
	return h
}

func (h *Handler[acc]) Clone() handler.Handler[acc] {
	return h
}

func (h *Handler[acc]) WithConfig(registry registry.Registry) handler.Handler[acc] {
	return h
}

func (h *Handler[acc]) Process(request gateway.Request,
	sess *session.Session[acc]) (response.Response, errors.Error) {
	return nil, errors.NotFound()
}

func (h *Handler[acc]) CanStepIn(sess *session.Session[acc]) bool {
	return false
}

func (h *Handler[acc]) CanStepOut(sess *session.Session[acc]) bool {
	return false
}

func (h *Handler[acc]) IsDone(sess *session.Session[acc]) bool {
	return false
}
