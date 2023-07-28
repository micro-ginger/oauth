package session

import (
	"encoding/json"

	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/session/domain/flow"
	"github.com/micro-ginger/oauth/login/session/domain/info"
)

type Session[acc account.Model] struct {
	Key       any
	Challenge string
	Flow      flow.Flow
	Info      *info.Info[acc]

	// state states of session
	// 0 new
	// 1 old
	state State
}

func (s *Session[acc]) MarshalBinary() (data []byte, err error) {
	var bytes []byte
	bytes, err = json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *Session[acc]) AddState(state State) {
	s.state.Add(state)
}

func (s *Session[acc]) IsFromDB() bool {
	return s.state.Has(StateFromDB)
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
