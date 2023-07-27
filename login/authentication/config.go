package authentication

import "github.com/micro-ginger/oauth/login/authentication/step"

type config struct {
	Steps map[step.Type]struct {
		Type step.Type
	}
}
