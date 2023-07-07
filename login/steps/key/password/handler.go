package password

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/steps/key"
)

type Handler[T account.Model] interface {
	step.Handler
	Initialize(getAccountHandleFunc key.GetAccountHandlerFunc[T])
}

type handler[T account.Model] struct {
	logger log.Logger

	getAccountHandleFunc key.GetAccountHandlerFunc[T]
}

func New[T account.Model](logger log.Logger) Handler[T] {
	h := &handler[T]{
		logger: logger,
	}
	return h
}

func (h *handler[T]) Initialize(getAccountHandleFunc key.GetAccountHandlerFunc[T]) {
	h.getAccountHandleFunc = getAccountHandleFunc
}
