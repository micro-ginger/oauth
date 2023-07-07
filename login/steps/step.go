package steps

import (
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/steps/key/password"
)

func GetByType(t step.Type) *step.Step {
	switch t {
	case password.Type:
		return password.Step
	}
	return nil
}
