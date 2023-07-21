package handler

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/authentication/handler/handler"
)

type Module[acc account.Model] struct {
	Handlers map[string]handler.Handler[acc]
}

func New[acc account.Model](logger log.Logger) *Module[acc] {
	m := &Module[acc]{
		Handlers: make(map[string]handler.Handler[acc]),
	}
	return m
}
