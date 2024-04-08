package flow

import (
	"github.com/ginger-core/compound/registry"
	"github.com/micro-ginger/oauth/login/flow/stage"
)

type Flows map[Section]*Flow

func (f Flows) Get(s Section, stg int) *Flow {
	flw := f[s]
	if flw == nil {
		return nil
	}
	if stg >= len(flw.Stages) {
		return nil
	}
	return &Flow{
		Section: flw.Section,
		Stages:  []stage.Stage{flw.Stages[stg]},
		Login:   flw.Login,
	}
}

func NewFlows(registry registry.Registry) Flows {
	flowsArr := make(Flows, 0)
	if err := registry.Unmarshal(&flowsArr); err != nil {
		panic(err)
	}
	flows := make(Flows)
	for _, flow := range flowsArr {
		flows[flow.Section] = flow
	}
	return flows
}
