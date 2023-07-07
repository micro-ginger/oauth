package key

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
)

type Handler[T account.Model] interface {
	step.Handler
	Initialize(getAccountHandleFunc GetAccountHandlerFunc[T])
}

type handler[T account.Model] struct {
	logger log.Logger

	getAccountHandleFunc GetAccountHandlerFunc[T]
}

func New[T account.Model](logger log.Logger) Handler[T] {
	h := &handler[T]{
		logger: logger,
	}
	return h
}

func (h *handler[T]) Initialize(getAccountHandleFunc GetAccountHandlerFunc[T]) {
	h.getAccountHandleFunc = getAccountHandleFunc
}
