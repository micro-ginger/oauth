package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	ld "github.com/micro-ginger/oauth/login/domain/delivery/login"
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
		// login
		sessions := make([]*session.CreateRequest, len(sess.Flow.Login.Sessions))
		for i, s := range sess.Flow.Login.Sessions {
			sessions[i] = new(session.CreateRequest)
			sessions[i].CreateConfig = s
			// add requested roles
			if len(sess.Info.RequestedRoles) > 0 {
				if sessions[i].CreateConfig.IncludeRoles == nil {
					sessions[i].CreateConfig.IncludeRoles = make([]string, 0)
				}
				sessions[i].CreateConfig.IncludeRoles =
					append(sessions[i].CreateConfig.IncludeRoles,
						sess.Info.RequestedRoles...)
			}
		}

		resp := ld.Response{
			Sessions: make(map[string]*ld.Session),
		}

		for _, s := range sessions {
			session, err := h.session.Create(request.GetContext(), s)
			if err != nil {
				return nil, err.WithTrace("session.Create")
			}
			resp.Sessions[session.Key] =
				ld.NewSession(session)
		}

		h.Respond(request, gateway.StatusOK, resp)
		return nil, nil
	}
	return r, err
}
