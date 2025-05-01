package steps

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/flow"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
)

type Module[acc account.Model, SessionAccountDetail gateway.ResultGetter] struct {
	logger log.Logger

	Handlers map[step.Type]handler.Handler[acc, SessionAccountDetail]
}

func New[acc account.Model, SessionAccountDetail gateway.ResultGetter](
	logger log.Logger,
) *Module[acc, SessionAccountDetail] {
	m := &Module[acc, SessionAccountDetail]{
		logger:   logger,
		Handlers: make(map[step.Type]handler.Handler[acc, SessionAccountDetail]),
	}
	return m
}

func (m *Module[acc, SessionAccountDetail]) Initialize(f flow.Flows) {
	for _, flow := range f {
		for _, stg := range flow.Stages {
			for si, step := range stg.Steps {
				if stp := GetByType(step.Type); stp != nil {
					stg.Steps[si].Populate(stp)
				}
			}
		}
	}
}

func (m *Module[acc, SessionAccountDetail]) RegisterHandler(
	hType step.Type, h handler.Handler[acc, SessionAccountDetail],
) {
	m.Handlers[hType] = h
}
