package refresh

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/authentication/handler/base"
	"github.com/micro-ginger/oauth/authentication/handler/handler"
	"github.com/micro-ginger/oauth/authentication/handler/mobile/account"
	"github.com/micro-ginger/oauth/authentication/info"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type h[acc account.Model] struct {
	*base.Handler[acc]

	logger log.Logger

	session session.UseCase
}

func New[acc account.Model](logger log.Logger,
	base *base.Handler[acc], session session.UseCase) handler.Handler[acc] {
	h := &h[acc]{
		Handler: base,
		logger:  logger,
		session: session,
	}
	return h
}

func (h *h[acc]) CanStepIn(info *info.Info[acc]) bool {
	return false
}

func (h *h[acc]) CanStepOut(info *info.Info[acc]) bool {
	return false
}

func (h *h[acc]) IsDone(info *info.Info[acc]) bool {
	return true
}
