package verify

import (
	"github.com/micro-ginger/oauth/login/authentication/step"
	"github.com/micro-ginger/oauth/login/authentication/step/action"
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
