package otp

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/login/authentication/info"
	"github.com/micro-ginger/oauth/login/authentication/otp"
	"github.com/micro-ginger/oauth/login/authentication/step"
	"github.com/micro-ginger/oauth/login/authentication/steps/base"
	"github.com/micro-ginger/oauth/login/authentication/steps/mobile/account"
	"github.com/micro-ginger/oauth/login/authentication/steps/mobile/verify"
)

type _handler[acc account.Model] struct {
	step.Handler[acc]
	logger log.Logger
	Base   *base.Handler[acc]
}

func New[acc account.Model](logger log.Logger, registry registry.Registry,
	base *base.Handler[acc], otp otp.Handler) step.Handler[acc] {
	h := &_handler[acc]{
		logger:  logger,
		Base:    base,
		Handler: verify.New(logger, registry, base, otp),
	}
	return h
}

func (h *_handler[acc]) Clone() step.Handler[acc] {
	return &_handler[acc]{
		logger:  h.logger,
		Base:    h.Base,
		Handler: h.Handler.Clone(),
	}
}

func (h *_handler[acc]) WithConfig(registry registry.Registry) step.Handler[acc] {
	h.Handler.WithConfig(registry)
	return h
}

func (h *_handler[acc]) CanStepIn(info *info.Info[acc]) bool {
	return info.StepInd != 0
}

func (h *_handler[acc]) CanStepOut(info *info.Info[acc]) bool {
	return info.StepInd > 1
}

func (h *_handler[acc]) IsDone(info *info.Info[acc]) bool {
	return info.StepInd > 1
}
