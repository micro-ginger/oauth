package login

import "github.com/micro-ginger/oauth/login/flow"

type Request struct {
	Section flow.Section `form:"section"`
	Stage   int          `form:"stage"`
}
