package otp

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/login/authentication/otp"
	"github.com/micro-ginger/oauth/login/authentication/steps/base"
	"github.com/micro-ginger/oauth/login/authentication/steps/mobile/account"
	"github.com/micro-ginger/oauth/login/authentication/steps/mobile/verify"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

type _handler[acc account.Model, SessionAccountDetail gateway.ResultGetter] struct {
	handler.Handler[acc, SessionAccountDetail]
	logger log.Logger
	Base   *base.Handler[acc, SessionAccountDetail]
}

func New[acc account.Model, SessionAccountDetail gateway.ResultGetter](logger log.Logger,
	registry registry.Registry, masker account.MaskerFunc,
	base *base.Handler[acc, SessionAccountDetail], otp otp.Handler) handler.Handler[acc, SessionAccountDetail] {
	h := &_handler[acc, SessionAccountDetail]{
		logger:  logger,
		Base:    base,
		Handler: verify.New(logger, registry, masker, base, otp),
	}
	return h
}

func (h *_handler[acc, SessionAccountDetail]) Clone() handler.Handler[acc, SessionAccountDetail] {
	return &_handler[acc, SessionAccountDetail]{
		logger:  h.logger,
		Base:    h.Base,
		Handler: h.Handler.Clone(),
	}
}

func (h *_handler[acc, SessionAccountDetail]) WithConfig(registry registry.Registry) handler.Handler[acc, SessionAccountDetail] {
	h.Handler.WithConfig(registry)
	return h
}

func (h *_handler[acc, SessionAccountDetail]) CanStepIn(sess *session.Session[acc, SessionAccountDetail]) bool {
	return sess.Flow.Pos.StepIndex != 0
}

func (h *_handler[acc, SessionAccountDetail]) CanStepOut(sess *session.Session[acc, SessionAccountDetail]) bool {
	return sess.Flow.Pos.StepIndex > 1
}

func (h *_handler[acc, SessionAccountDetail]) IsDone(sess *session.Session[acc, SessionAccountDetail]) bool {
	return sess.Flow.Pos.StepIndex > 1
}
