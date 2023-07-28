package authentication

import "github.com/micro-ginger/oauth/login/flow/stage/step"

type config struct {
	Steps map[string]struct {
		Type step.Type
	}
}
