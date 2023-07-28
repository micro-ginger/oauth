package flow

import (
	"github.com/ginger-core/compound/registry"
)

type Flows map[Section]*Flow

func (f Flows) Get(s Section) *Flow {
	return f[s]
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
