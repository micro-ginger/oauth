package verify

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/login/authentication/info"
	"github.com/micro-ginger/oauth/login/authentication/otp"
	"github.com/micro-ginger/oauth/login/authentication/step"
	"github.com/micro-ginger/oauth/login/authentication/steps/base"
	"github.com/micro-ginger/oauth/login/authentication/steps/mobile/account"
)

type _handler[acc account.Model] struct {
	*base.Handler[acc]

	logger log.Logger
	config config

	otp otp.Handler
}

func New[acc account.Model](logger log.Logger, registry registry.Registry,
	base *base.Handler[acc], otp otp.Handler) step.Handler[acc] {
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

func (h *_handler[acc]) Clone() step.Handler[acc] {
	return &_handler[acc]{
		Handler: h.Handler,
		logger:  h.logger,
		otp:     h.otp,
	}
}

func (h *_handler[acc]) WithConfig(registry registry.Registry) step.Handler[acc] {
	if err := registry.Unmarshal(&h.config); err != nil {
		panic(err)
	}
	return h
}

func (h *_handler[acc]) CanStepIn(info *info.Info[acc]) bool {
	return info.StepInd == 0
}

func (h *_handler[acc]) CanStepOut(info *info.Info[acc]) bool {
	return info.StepInd == 1
}

func (h *_handler[acc]) IsDone(info *info.Info[acc]) bool {
	return info.StepInd > 1
}