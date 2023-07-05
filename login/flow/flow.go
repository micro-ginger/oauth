package flow

import "github.com/micro-ginger/oauth/login/flow/stage"

type Flow struct {
	Section Section
	Stages  []stage.Stage
	Login   Login
}
