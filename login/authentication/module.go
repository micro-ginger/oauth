package authentication

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	loginSession "github.com/micro-ginger/oauth/login/session/domain/session"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Model[acc a.ExtentedModel] interface {
	Initialize()
	GetStepHandlers() map[step.Type]handler.Handler[acc]
}

type Module[acc a.ExtentedModel] interface {
	Model[acc]
	GetBase() *Base[acc]
}

type module[acc a.ExtentedModel] struct {
	*Base[acc]
}

func New[acc a.ExtentedModel](
	logger log.Logger, registry registry.Registry,
	loginSession loginSession.Handler[acc],
	cache repository.Cache, account account.UseCase[acc],
	session session.UseCase[acc],
) Module[acc] {
	m := &module[acc]{
		Base: NewBase(logger, registry,
			loginSession, cache, account, session),
	}

	return m
}
