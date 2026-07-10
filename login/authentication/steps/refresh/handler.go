package refresh

import (
	"github.com/ginger-core/log"
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/login/authentication/steps/base"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	s "github.com/micro-ginger/oauth/login/session/domain/session"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type h[acc a.ExtentedModel] struct {
	*base.Handler[acc]

	logger log.Logger

	session session.UseCase[acc]
}

func New[acc a.ExtentedModel](logger log.Logger,
	base *base.Handler[acc],
	session session.UseCase[acc],
) handler.Handler[acc] {
	h := &h[acc]{
		Handler: base,
		logger:  logger,
		session: session,
	}
	return h
}

func (h *h[acc]) CanStepIn(
	sess *s.Session[acc]) bool {
	return false
}

func (h *h[acc]) CanStepOut(
	sess *s.Session[acc]) bool {
	return false
}

func (h *h[acc]) IsDone(
	sess *s.Session[acc]) bool {
	return true
}
