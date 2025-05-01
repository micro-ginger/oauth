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

type Handler[acc account.Model, SessionAccountDetail gateway.ResultGetter] struct {
	logger log.Logger

	Type step.Type

	Account account.UseCase[acc]
	Session session.Handler[acc, SessionAccountDetail]

	AccountGetter info.AccountGetter[acc]

	SuspendValidator validator.UseCase
}

func New[acc account.Model, SessionAccountDetail gateway.ResultGetter](
	logger log.Logger, registry registry.Registry,
	session session.Handler[acc, SessionAccountDetail],
	cache repository.Cache) *Handler[acc, SessionAccountDetail] {
	// suspendValidator is otp validation which is being
	// validated in each login session
	suspendValidator := v.New(
		logger.WithTrace("validators.suspend"),
		registry.ValueOf("validators.suspend"),
		cache,
	)

	h := &Handler[acc, SessionAccountDetail]{
		logger:           logger,
		Session:          session,
		SuspendValidator: suspendValidator.UseCase,
	}
	return h
}

func (h *Handler[acc, SessionAccountDetail]) WithType(t step.Type) handler.Handler[acc, SessionAccountDetail] {
	h.Type = t
	return h
}

func (h *Handler[acc, SessionAccountDetail]) GetType() step.Type {
	return h.Type
}

func (h *Handler[acc, SessionAccountDetail]) WithAccount(
	account account.UseCase[acc]) handler.Handler[acc, SessionAccountDetail] {
	h.Account = account
	return h
}

func (h *Handler[acc, SessionAccountDetail]) WithAccountGetter(
	getter info.AccountGetter[acc]) handler.Handler[acc, SessionAccountDetail] {
	h.AccountGetter = getter
	return h
}

func (h *Handler[acc, SessionAccountDetail]) Clone() handler.Handler[acc, SessionAccountDetail] {
	return h
}

func (h *Handler[acc, SessionAccountDetail]) WithConfig(registry registry.Registry) handler.Handler[acc, SessionAccountDetail] {
	return h
}

func (h *Handler[acc, SessionAccountDetail]) Process(request gateway.Request,
	sess *session.Session[acc, SessionAccountDetail]) (response.Response, errors.Error) {
	return nil, errors.NotFound()
}

func (h *Handler[acc, SessionAccountDetail]) CanStepIn(sess *session.Session[acc, SessionAccountDetail]) bool {
	return false
}

func (h *Handler[acc, SessionAccountDetail]) CanStepOut(sess *session.Session[acc, SessionAccountDetail]) bool {
	return false
}

func (h *Handler[acc, SessionAccountDetail]) IsDone(sess *session.Session[acc, SessionAccountDetail]) bool {
	return false
}
