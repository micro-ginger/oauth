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

type Handler[acc account.Model] interface {
	WithType(t step.Type) Handler[acc]
	GetType() step.Type

	Clone() Handler[acc]
	WithConfig(registry registry.Registry) Handler[acc]
	WithAccount(account account.UseCase[acc]) Handler[acc]
	WithAccountGetter(getter info.AccountGetter[acc]) Handler[acc]

	Process(request gateway.Request,
		sess *session.Session[acc]) (response.Response, errors.Error)

	CanStepIn(sess *session.Session[acc]) bool
	CanStepOut(sess *session.Session[acc]) bool
	IsDone(sess *session.Session[acc]) bool
}
