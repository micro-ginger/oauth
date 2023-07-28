package otp

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/login/authentication/otp"
	"github.com/micro-ginger/oauth/login/authentication/steps/base"
	"github.com/micro-ginger/oauth/login/authentication/steps/mobile/account"
	"github.com/micro-ginger/oauth/login/authentication/steps/mobile/verify"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

type _handler[acc account.Model] struct {
	handler.Handler[acc]
	logger log.Logger
	Base   *base.Handler[acc]
}

func New[acc account.Model](logger log.Logger, registry registry.Registry,
	base *base.Handler[acc], otp otp.Handler) handler.Handler[acc] {
	h := &_handler[acc]{
		logger:  logger,
		Base:    base,
		Handler: verify.New(logger, registry, base, otp),
	}
	return h
}

func (h *_handler[acc]) Clone() handler.Handler[acc] {
	return &_handler[acc]{
		logger:  h.logger,
		Base:    h.Base,
		Handler: h.Handler.Clone(),
	}
}

func (h *_handler[acc]) WithConfig(registry registry.Registry) handler.Handler[acc] {
	h.Handler.WithConfig(registry)
	return h
}

func (h *_handler[acc]) CanStepIn(sess *session.Session[acc]) bool {
	return sess.Flow.Pos.StepIndex != 0
}

func (h *_handler[acc]) CanStepOut(sess *session.Session[acc]) bool {
	return sess.Flow.Pos.StepIndex > 1
}

func (h *_handler[acc]) IsDone(sess *session.Session[acc]) bool {
	return sess.Flow.Pos.StepIndex > 1
}
