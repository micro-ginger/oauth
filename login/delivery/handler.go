package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authorization"
	ad "github.com/micro-ginger/oauth/login/domain/account"
	ld "github.com/micro-ginger/oauth/login/domain/delivery/login"
	"github.com/micro-ginger/oauth/login/flow"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	s "github.com/micro-ginger/oauth/login/session/domain/session"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Handler[acc account.Model, SessionAccountDetail gateway.ResultGetter] interface {
	gateway.Handler
	Initialize(account ad.UseCase[acc],
		loginSession s.Handler[acc, SessionAccountDetail], flows flow.Flows,
		session session.UseCase[SessionAccountDetail],
	)
	SetManager(manager authorization.Manager[acc, SessionAccountDetail])
	RegisterHandler(t step.Type, sh handler.Handler[acc, SessionAccountDetail])
}

type lh[acc account.Model, SessionAccountDetail gateway.ResultGetter] struct {
	gateway.Responder
	logger log.Logger

	account ad.UseCase[acc]

	loginSession s.Handler[acc, SessionAccountDetail]

	flows   flow.Flows
	session session.UseCase[SessionAccountDetail]

	stepHandlers map[step.Type]handler.Handler[acc, SessionAccountDetail]

	manager authorization.Manager[acc, SessionAccountDetail]
}

func NewLogin[acc account.Model, SessionAccountDetail gateway.ResultGetter](
	logger log.Logger, responder gateway.Responder,
) Handler[acc, SessionAccountDetail] {
	h := &lh[acc, SessionAccountDetail]{
		Responder: responder,
		logger:    logger,
	}
	return h
}

func (h *lh[acc, SessionAccountDetail]) Initialize(
	account ad.UseCase[acc], loginSession s.Handler[acc, SessionAccountDetail],
	flows flow.Flows, session session.UseCase[SessionAccountDetail]) {
	h.account = account
	h.loginSession = loginSession
	h.flows = flows
	h.session = session
}

func (h *lh[acc, SessionAccountDetail]) SetManager(manager authorization.Manager[acc, SessionAccountDetail]) {
	h.manager = manager
}

func (h *lh[acc, SessionAccountDetail]) RegisterHandler(
	t step.Type, sh handler.Handler[acc, SessionAccountDetail],
) {
	if h.stepHandlers == nil {
		h.stepHandlers = make(map[step.Type]handler.Handler[acc, SessionAccountDetail])
	}
	h.stepHandlers[t] = sh
}

func (h *lh[acc, SessionAccountDetail]) Handle(request gateway.Request) (r any, err errors.Error) {
	if h.manager != nil {
		// before handle request
		if err = h.manager.BeforeHandle(request); err != nil {
			return nil, err.WithTrace("manager.BeforeHandle")
		}
	}
	var sess *s.Session[acc, SessionAccountDetail]
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
	if err = h.validate(request, sess); err != nil {
		return nil, err.WithTrace("validate")
	}
	if sess.IsDone() {
		if h.manager != nil {
			// before login
			if err = h.manager.BeforeLogin(request, sess); err != nil {
				return nil, err.WithTrace("manager.BeforeLogin")
			}
		}
		if err = h.validate(request, sess); err != nil {
			return nil, err.WithTrace("validate")
		}
		// login
		sessions := make([]*session.CreateRequest[SessionAccountDetail], len(sess.Flow.Login.Sessions))
		for i, s := range sess.Flow.Login.Sessions {
			sessions[i] = new(session.CreateRequest[SessionAccountDetail])
			sessions[i].CreateConfig = s
			sessions[i].CreateConfig.Section = sess.Info.Section
			// populate account
			sessions[i].Account.Id = sess.Info.AccountId
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
		if s := sess.Info.GetTemp("session"); s != nil {
			sess := s.(*session.Session[SessionAccountDetail])
			sessions = append(sessions,
				&session.CreateRequest[SessionAccountDetail]{
					CreateConfig: session.NewCreateConfigFromSession(sess),
					Old:          sess,
				},
			)
		}

		resp := &ld.Response[SessionAccountDetail]{
			Sessions: make(map[string]*ld.Session[SessionAccountDetail]),
		}

		if h.manager != nil {
			// before session create
			if err = h.manager.BeforeSessionCreate(
				request, sess, sessions); err != nil {
				return nil, err.
					WithTrace("manager.BeforeSessionCreate")
			}
		}
		for _, s := range sessions {
			session, err := h.session.Create(request.GetContext(), s)
			if err != nil {
				return nil, err.WithTrace("session.Create")
			}
			resp.Sessions[session.Section] = ld.NewSession(session)
		}
		if h.manager != nil {
			// after session create
			if err = h.manager.AfterSessionCreate(
				request, sess, resp); err != nil {
				return nil, err.
					WithTrace("manager.AfterSessionCreate")
			}
		}

		h.Respond(request, gateway.StatusOK, resp)
		return nil, nil
	}
	return r, err
}
