package authentication

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	loginSession "github.com/micro-ginger/oauth/login/session/domain/session"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Model[acc account.Model, SessionAccountDetail gateway.ResultGetter] interface {
	Initialize()
	GetStepHandlers() map[step.Type]handler.Handler[acc, SessionAccountDetail]
}

type Module[acc account.Model, SessionAccountDetail gateway.ResultGetter] interface {
	Model[acc, SessionAccountDetail]
	GetBase() *Base[acc, SessionAccountDetail]
}

type module[acc account.Model, SessionAccountDetail gateway.ResultGetter] struct {
	*Base[acc, SessionAccountDetail]
}

func New[acc account.Model, SessionAccountDetail gateway.ResultGetter](
	logger log.Logger, registry registry.Registry,
	loginSession loginSession.Handler[acc, SessionAccountDetail],
	cache repository.Cache, account account.UseCase[acc],
	session session.UseCase[SessionAccountDetail],
) Module[acc, SessionAccountDetail] {
	m := &module[acc, SessionAccountDetail]{
		Base: NewBase(logger, registry,
			loginSession, cache, account, session),
	}

	return m
}
