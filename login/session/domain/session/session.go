package session

import (
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/session/domain/flow"
	"github.com/micro-ginger/oauth/login/session/domain/info"
)

type Session[acc account.Model] struct {
	Key       any
	Challenge string
	Flow      flow.Flow
	Info      *info.Info[acc]
}

func (s *Session[acc]) GetKey() string {
	return "login.sessions." + s.Challenge
}

func (s *Session[acc]) Next() {
	stage := s.Flow.Stages[s.Flow.Pos.StageIndex]
	step := stage.Steps[s.Flow.Pos.StepIndex]
	if s.Flow.Pos.ActionIndex+1 >= len(step.Actions) {
		// next stage
		s.Flow.Pos.ActionIndex = 0
		s.Flow.Pos.StageIndex++
	} else {
		s.Flow.Pos.ActionIndex++
	}
}

func (s *Session[acc]) IsDone() bool {
	return s.Flow.Pos.StageIndex >= len(s.Flow.Stages)
}
