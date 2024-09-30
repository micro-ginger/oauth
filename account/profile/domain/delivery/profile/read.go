package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-blonde/auth/profile"
	"github.com/micro-blonde/file"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type Read[T profile.Model, F file.Model] interface {
	PopulateRead(request gateway.Request, prof *p.Profile[T]) errors.Error
}
