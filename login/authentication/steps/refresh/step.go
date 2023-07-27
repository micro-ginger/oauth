package refresh

import (
	"github.com/micro-ginger/oauth/login/authentication/step"
	"github.com/micro-ginger/oauth/login/authentication/step/action"
)

const Type step.Type = "REFRESH"

var (
	Step = &step.Step{
		Type: Type,
		Actions: []action.Action{
			{
				Type: "VERIFY",
			},
		},
	}
)
