package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/flow"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	s "github.com/micro-ginger/oauth/login/session/domain/session"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Handler[acc account.Model] interface {
	gateway.Handler
	Initialize(loginSession s.Handler[acc],
		flows flow.Flows, session session.UseCase)
	RegisterHandler(t step.Type, sh handler.Handler[acc])
}

type lh[acc account.Model] struct {
	gateway.Responder
	logger log.Logger

	loginSession s.Handler[acc]

	flows   flow.Flows
	session session.UseCase

	stepHandlers map[step.Type]handler.Handler[acc]
}

func NewLogin[acc account.Model](
	logger log.Logger, responder gateway.Responder) Handler[acc] {
	h := &lh[acc]{
		Responder: responder,
		logger:    logger,
	}
	return h
}

func (h *lh[acc]) Initialize(loginSession s.Handler[acc],
	flows flow.Flows, session session.UseCase) {
	h.loginSession = loginSession
	h.flows = flows
	h.session = session
}

func (h *lh[acc]) RegisterHandler(t step.Type, sh handler.Handler[acc]) {
	if h.stepHandlers == nil {
		h.stepHandlers = make(map[step.Type]handler.Handler[acc])
	}
	h.stepHandlers[t] = sh
}

func (h *lh[acc]) Handle(request gateway.Request) (r any, err errors.Error) {
	var sess *s.Session[acc]
	challenge, ok := request.GetQuery("challenge")
	if ok {
		sess, r, err = h.challenge(request, challenge)
		if err != nil {
			return nil, err.WithTrace("challenge")
		}
	} else {
		sess, r, err = h.start(request)
		if err != nil {
			return nil, err.WithTrace("start")
		}
	}
	if sess.IsDone() {
		// TODO login
		return nil, nil
	}
	return r, err
}
