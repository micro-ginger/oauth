package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/login/flow"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
)

type lh struct {
	gateway.Responder
	logger log.Logger

	flows flow.Flows

	stepHandlers map[step.Type]step.Handler
}

func NewLogin(logger log.Logger, responder gateway.Responder) gateway.Handler {
	h := &lh{
		Responder: responder,
		logger:    logger,
	}
	return h
}

func (h *lh) RegisterHandler(t step.Type, sh step.Handler) {
	if h.stepHandlers == nil {
		h.stepHandlers = make(map[step.Type]step.Handler)
	}
	h.stepHandlers[t] = sh
}

func (h *lh) Handle(request gateway.Request) (any, errors.Error) {
	challenge, ok := request.GetQuery("challenge")
	if ok {
		return h.challenge(request, challenge)
	}
	return h.start(request)
}
