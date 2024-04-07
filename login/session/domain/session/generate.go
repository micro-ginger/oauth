package session

import "github.com/micro-ginger/oauth/login/flow"

type GenerateRequest struct {
	Flow  *flow.Flow
	Step  string
	Roles []string
}
