package handler

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication/response"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/session/domain/info"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

type Handler[acc account.Model, SessionAccountDetail gateway.ResultGetter] interface {
	WithType(t step.Type) Handler[acc, SessionAccountDetail]
	GetType() step.Type

	Clone() Handler[acc, SessionAccountDetail]
	WithConfig(registry registry.Registry) Handler[acc, SessionAccountDetail]
	WithAccount(account account.UseCase[acc]) Handler[acc, SessionAccountDetail]
	WithAccountGetter(getter info.AccountGetter[acc]) Handler[acc, SessionAccountDetail]

	Process(request gateway.Request,
		sess *session.Session[acc, SessionAccountDetail],
	) (response.Response, errors.Error)

	CanStepIn(sess *session.Session[acc, SessionAccountDetail]) bool
	CanStepOut(sess *session.Session[acc, SessionAccountDetail]) bool
	IsDone(sess *session.Session[acc, SessionAccountDetail]) bool
}
