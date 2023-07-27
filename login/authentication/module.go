package authentication

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Module interface {
	Initialize()
}

type module struct {
	Module
}

func New[acc account.Model](logger log.Logger, registry registry.Registry,
	account account.UseCase[acc], session session.UseCase) Module {
	m := &module{
		Module: newBase(logger, registry, account, session),
	}

	return m
}
