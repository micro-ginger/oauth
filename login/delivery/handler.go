package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication/step"
	"github.com/micro-ginger/oauth/login/flow"
	"github.com/micro-ginger/oauth/login/session"
)

type Handler[acc account.Model] interface {
	gateway.Handler
	Initialize(flows flow.Flows, session session.UseCase)
	RegisterHandler(t step.Type, sh step.Handler[acc])
}

type lh[acc account.Model] struct {
	gateway.Responder
	logger log.Logger

	flows   flow.Flows
	session session.UseCase

	stepHandlers map[step.Type]step.Handler[acc]
}

func NewLogin[acc account.Model](
	logger log.Logger, responder gateway.Responder) Handler[acc] {
	h := &lh[acc]{
		Responder: responder,
		logger:    logger,
	}
	return h
}

func (h *lh[acc]) Initialize(flows flow.Flows, session session.UseCase) {
	h.flows = flows
	h.session = session
}

func (h *lh[acc]) RegisterHandler(t step.Type, sh step.Handler[acc]) {
	if h.stepHandlers == nil {
		h.stepHandlers = make(map[step.Type]step.Handler[acc])
	}
	h.stepHandlers[t] = sh
}

func (h *lh[acc]) Handle(request gateway.Request) (r any, err errors.Error) {
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
