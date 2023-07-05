package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/domain/login"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
)

func (h *lh) process(request gateway.Request,
	session *login.Session) (any, errors.Error) {
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
	// next session pos
	// store session
	return r, nil
}
