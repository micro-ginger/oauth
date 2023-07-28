package verify

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/login/authentication/otp"
	"github.com/micro-ginger/oauth/login/authentication/steps/base"
	"github.com/micro-ginger/oauth/login/authentication/steps/mobile/account"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

type _handler[acc account.Model] struct {
	*base.Handler[acc]

	logger log.Logger
	config config

	otp otp.Handler
}

func New[acc account.Model](logger log.Logger, registry registry.Registry,
	base *base.Handler[acc], otp otp.Handler) handler.Handler[acc] {
	h := &_handler[acc]{
		Handler: base,
		logger:  logger,
		otp:     otp,
	}
	if err := registry.Unmarshal(&h.config); err != nil {
		panic(err)
	}
	h.config.initialize()
	return h
}

func (h *_handler[acc]) Clone() handler.Handler[acc] {
	return &_handler[acc]{
		Handler: h.Handler,
		logger:  h.logger,
		otp:     h.otp,
	}
}

func (h *_handler[acc]) WithConfig(registry registry.Registry) handler.Handler[acc] {
	if err := registry.Unmarshal(&h.config); err != nil {
		panic(err)
	}
	return h
}

func (h *_handler[acc]) CanStepIn(sess *session.Session[acc]) bool {
	return sess.Flow.Pos.StepIndex == 0
}

func (h *_handler[acc]) CanStepOut(sess *session.Session[acc]) bool {
	return sess.Flow.Pos.StepIndex == 1
}

func (h *_handler[acc]) IsDone(sess *session.Session[acc]) bool {
	return sess.Flow.Pos.StepIndex > 1
}
