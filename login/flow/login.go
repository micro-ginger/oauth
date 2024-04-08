package flow

import "github.com/micro-ginger/oauth/session/domain/session"

type Login struct {
	Sessions []*session.CreateConfig
}
