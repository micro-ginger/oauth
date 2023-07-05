package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/login/flow"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/session"
)

type Login interface {
	gateway.Handler
	Initialize(flows flow.Flows, session session.UseCase)
}

type lh struct {
	gateway.Responder
	logger log.Logger

	flows   flow.Flows
	session session.UseCase

	stepHandlers map[step.Type]step.Handler
}

func NewLogin(logger log.Logger, responder gateway.Responder) Login {
	h := &lh{
		Responder: responder,
		logger:    logger,
	}
	return h
}

func (h *lh) Initialize(flows flow.Flows, session session.UseCase) {
	h.session = session
}

func (h *lh) RegisterHandler(t step.Type, sh step.Handler) {
	if h.stepHandlers == nil {
		h.stepHandlers = make(map[step.Type]step.Handler)
	}
	h.stepHandlers[t] = sh
}

func (h *lh) Handle(request gateway.Request) (r any, err errors.Error) {
	var session *session.Session
	challenge, ok := request.GetQuery("challenge")
	if ok {
		session, r, err = h.challenge(request, challenge)
		if err != nil {
			return nil, err.WithTrace("challenge")
		}
	} else {
		session, r, err = h.start(request)
		if err != nil {
			return nil, err.WithTrace("start")
		}
	}
	if session.IsDone() {
		// TODO login
		return nil, nil
	}
	return r, err
}
