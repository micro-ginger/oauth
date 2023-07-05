package stage

import "github.com/micro-ginger/oauth/login/flow/stage/step"

type Stage struct {
	Steps []*step.Step
	Next  *Stage
}

func (m *Stage) GetStepIndex(t step.Type) int {
	for i, s := range m.Steps {
		if s.Type == t {
			return i
		}
	}
	return -1
}
