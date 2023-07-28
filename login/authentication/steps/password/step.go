package password

import (
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/action"
)

const Type step.Type = "PASSWORD"

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
