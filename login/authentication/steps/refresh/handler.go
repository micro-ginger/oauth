package refresh

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication/steps/base"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	s "github.com/micro-ginger/oauth/login/session/domain/session"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type h[acc account.Model, SessionAccountDetail gateway.ResultGetter] struct {
	*base.Handler[acc, SessionAccountDetail]

	logger log.Logger

	session session.UseCase[SessionAccountDetail]
}

func New[acc account.Model, SessionAccountDetail gateway.ResultGetter](logger log.Logger,
	base *base.Handler[acc, SessionAccountDetail],
	session session.UseCase[SessionAccountDetail],
) handler.Handler[acc, SessionAccountDetail] {
	h := &h[acc, SessionAccountDetail]{
		Handler: base,
		logger:  logger,
		session: session,
	}
	return h
}

func (h *h[acc, SessionAccountDetail]) CanStepIn(
	sess *s.Session[acc, SessionAccountDetail]) bool {
	return false
}

func (h *h[acc, SessionAccountDetail]) CanStepOut(
	sess *s.Session[acc, SessionAccountDetail]) bool {
	return false
}

func (h *h[acc, SessionAccountDetail]) IsDone(
	sess *s.Session[acc, SessionAccountDetail]) bool {
	return true
}
