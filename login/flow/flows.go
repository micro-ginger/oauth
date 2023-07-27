package flow

import (
	"github.com/ginger-core/compound/registry"
	"github.com/micro-ginger/oauth/login/authentication/steps"
)

type Flows map[Section]*Flow

func (f Flows) Get(s Section) *Flow {
	return f[s]
}

func (f Flows) Initialize() {
	for _, flow := range f {
		for _, stg := range flow.Stages {
			for si, step := range stg.Steps {
				if stp := steps.GetByType(step.Type); stp != nil {
					stg.Steps[si] = stp
				}
			}
		}
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
