package verify

import (
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/action"
)

const Type step.Type = "MOBILE_VERIFY"

var (
	Step = &step.Step{
		Type: Type,
		Actions: []action.Action{
			{
				Type: "REQUEST",
			},
			{
				Type: "VERIFY",
			},
		},
	}
)
