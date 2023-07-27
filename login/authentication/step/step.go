package step

import "github.com/micro-ginger/oauth/login/authentication/step/action"

type Step struct {
	Type    Type
	Actions []action.Action
}
