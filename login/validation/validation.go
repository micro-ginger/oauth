package validation

import (
	"github.com/micro-blonde/auth/account"
)

type Validation struct {
	Step     Step
	Type     Type
	Status   account.Status
	ErrorKey string
}
