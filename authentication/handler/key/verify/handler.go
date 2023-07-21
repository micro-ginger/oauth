package verify

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/authentication/handler/base"
	"github.com/micro-ginger/oauth/authentication/handler/handler"
	"github.com/micro-ginger/oauth/authentication/info"
)

type h[acc account.Model] struct {
	*base.Handler[acc]

	logger log.Logger
	config config
}

func New[acc account.Model](logger log.Logger, registry registry.Registry,
	base *base.Handler[acc]) handler.Handler[acc] {
	h := &h[acc]{
		Handler: base,
		logger:  logger,
	}
	if registry != nil {
		if err := registry.Unmarshal(&h.config); err != nil {
			panic(err)
		}
	}
	h.config.initialize()
	return h
}

func (h *h[acc]) CanStepIn(info *info.Info[acc]) bool {
	return false
}

func (h *h[acc]) CanStepOut(info *info.Info[acc]) bool {
	return true
}

func (h *h[acc]) IsDone(info *info.Info[acc]) bool {
	return true
}
