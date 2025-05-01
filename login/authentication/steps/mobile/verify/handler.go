package verify

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/login/authentication/otp"
	"github.com/micro-ginger/oauth/login/authentication/steps/base"
	"github.com/micro-ginger/oauth/login/authentication/steps/mobile/account"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

type _handler[acc account.Model, SessionAccountDetail gateway.ResultGetter] struct {
	*base.Handler[acc, SessionAccountDetail]

	logger log.Logger
	config config

	otp otp.Handler

	masker account.MaskerFunc
}

func New[acc account.Model, SessionAccountDetail gateway.ResultGetter](logger log.Logger,
	registry registry.Registry, masker account.MaskerFunc,
	base *base.Handler[acc, SessionAccountDetail], otp otp.Handler,
) handler.Handler[acc, SessionAccountDetail] {
	h := &_handler[acc, SessionAccountDetail]{
		Handler: base,
		logger:  logger,
		otp:     otp,
		masker:  masker,
	}
	if err := registry.Unmarshal(&h.config); err != nil {
		panic(err)
	}
	h.config.initialize()
	return h
}

func (h *_handler[acc, SessionAccountDetail]) Clone() handler.Handler[acc, SessionAccountDetail] {
	return &_handler[acc, SessionAccountDetail]{
		Handler: h.Handler,
		logger:  h.logger,
		otp:     h.otp,
	}
}

func (h *_handler[acc, SessionAccountDetail]) WithConfig(
	registry registry.Registry) handler.Handler[acc, SessionAccountDetail] {
	if err := registry.Unmarshal(&h.config); err != nil {
		panic(err)
	}
	return h
}

func (h *_handler[acc, SessionAccountDetail]) CanStepIn(
	sess *session.Session[acc, SessionAccountDetail]) bool {
	return sess.Flow.Pos.StepIndex == 0
}

func (h *_handler[acc, SessionAccountDetail]) CanStepOut(
	sess *session.Session[acc, SessionAccountDetail]) bool {
	return sess.Flow.Pos.StepIndex == 1
}

func (h *_handler[acc, SessionAccountDetail]) IsDone(
	sess *session.Session[acc, SessionAccountDetail]) bool {
	return sess.Flow.Pos.StepIndex > 1
}
