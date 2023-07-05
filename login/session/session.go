package session

type Session struct {
	challenge string
	Flow      Flow
}

func (s *Session) GetKey() string {
	return "login.sessions." + s.challenge
}

func (s *Session) Next() {
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

func (s *Session) IsDone() bool {
	return s.Flow.Pos.StageIndex >= len(s.Flow.Stages)
}
