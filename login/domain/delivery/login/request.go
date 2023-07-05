package login

import "github.com/micro-ginger/oauth/login/flow"

type Request struct {
	Section flow.Section `query:"section"`
}
