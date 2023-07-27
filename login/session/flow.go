package session

import (
	"github.com/micro-ginger/oauth/login/authentication/step"
	"github.com/micro-ginger/oauth/login/flow"
)

type Flow struct {
	*flow.Flow
	Pos Position
}

func (m *Flow) GetCurrentStep() (*step.Step, int) {
	stage := m.Flow.Stages[m.Pos.StageIndex]
	step := stage.Steps[m.Pos.StepIndex]
	return step, m.Pos.ActionIndex
}
