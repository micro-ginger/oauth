package steps

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/authentication/steps/handler"
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

func (m *Module[acc]) RegisterHandler(hType string, h handler.Handler[acc]) {
	m.Handlers[hType] = h
}
