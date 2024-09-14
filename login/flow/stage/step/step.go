package step

import "github.com/micro-ginger/oauth/login/flow/stage/step/action"

type Step struct {
	Type              Type
	Actions           []action.Action
	IsCaptchaRequired bool
}

func (s *Step) Populate(itm *Step) {
	s.Actions = itm.Actions
}
