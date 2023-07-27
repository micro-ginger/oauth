package steps

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication/step"
	"github.com/micro-ginger/oauth/login/authentication/steps/base"
)

type Module[acc account.Model] struct {
	logger log.Logger

	Handlers map[string]step.Handler[acc]
	base     *base.Handler[acc]
}

func New[acc account.Model](logger log.Logger) *Module[acc] {
	m := &Module[acc]{
		logger:   logger,
		Handlers: make(map[string]step.Handler[acc]),
	}
	return m
}

func (m *Module[acc]) RegisterHandler(hType string, h step.Handler[acc]) {
	m.Handlers[hType] = h
}
