package steps

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/flow"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
)

type Module[acc account.Model] struct {
	logger log.Logger

	Handlers map[step.Type]handler.Handler[acc]
}

func New[acc account.Model](logger log.Logger) *Module[acc] {
	m := &Module[acc]{
		logger:   logger,
		Handlers: make(map[step.Type]handler.Handler[acc]),
	}
	return m
}

func (m *Module[acc]) Initialize(f flow.Flows) {
	for _, flow := range f {
		for _, stg := range flow.Stages {
			for si, step := range stg.Steps {
				if stp := GetByType(step.Type); stp != nil {
					stg.Steps[si] = stp
				}
			}
		}
	}
}

func (m *Module[acc]) RegisterHandler(hType step.Type, h handler.Handler[acc]) {
	m.Handlers[hType] = h
}
