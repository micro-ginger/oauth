package step

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
)

type Meta struct {
	ActionIndex int
}

type Handler interface {
	Process(request gateway.Request, meta *Meta) (any, errors.Error)
}
