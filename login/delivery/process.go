package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/session"
)

func (h *lh) process(request gateway.Request,
	session *session.Session) (any, errors.Error) {
	s, actionIdx := session.Flow.GetCurrentStep()
	sh := h.stepHandlers[s.Type]

	r, err := sh.Process(request,
		&step.Meta{
			ActionIndex: actionIdx,
		})
	if err != nil {
		return nil, err.
			WithDetail(errors.NewDetail().
				With("step", s).
				With("actionIdx", actionIdx)).
			WithTrace("sh.Process")
	}
	// next
	session.Next()
	if session.IsDone() {
		err = h.session.Delete(request.GetContext(), session)
		if err != nil {
			return nil, err.WithTrace("session.Delete")
		}
		return nil, nil
	}
	// store session
	err = h.session.Store(request.GetContext(), session)
	if err != nil {
		return nil, err.WithTrace("session.Store")
	}
	return r, nil
}
