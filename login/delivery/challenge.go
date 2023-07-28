package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *lh[acc]) challenge(request gateway.Request,
	challenge string) (*session.Session[acc], any, errors.Error) {
	// TODO get process
	// process
	return nil, nil, nil
}
