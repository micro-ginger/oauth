package flow

import (
	"github.com/micro-ginger/oauth/login/validation"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Login struct {
	Validations []validation.Validation
	Sessions    []*session.CreateConfig
}
