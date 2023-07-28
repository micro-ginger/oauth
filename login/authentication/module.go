package authentication

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	loginSession "github.com/micro-ginger/oauth/login/session/domain/session"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Model[acc account.Model] interface {
	Initialize()
	GetStepHandlers() map[step.Type]handler.Handler[acc]
}

type Module[acc account.Model] struct {
	Model[acc]
}

func New[acc account.Model](logger log.Logger, registry registry.Registry,
	loginSession loginSession.Handler[acc], cache repository.Cache,
	account account.UseCase[acc], session session.UseCase) Model[acc] {
	m := &Module[acc]{
		Model: NewBase(logger, registry,
			loginSession, cache, account, session),
	}

	return m
}
