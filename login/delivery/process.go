package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *lh[acc]) process(request gateway.Request,
	sess *session.Session[acc]) (any, errors.Error) {
	s, actionIdx := sess.Flow.GetCurrentStep()
	sh := h.stepHandlers[s.Type]
	if sh == nil {
		return nil, errors.Unauthorized().
			WithTrace("sh=nil").
			WithDesc("step handler not found")
	}

	r, err := sh.Process(request, sess)
	if err != nil {
		return nil, err.
			WithDetail(errors.NewDetail().
				With("step", s).
				With("actionIdx", actionIdx)).
			WithTrace("sh.Process")
	}
	ctx := request.GetContext()
	// next
	sess.Next()
	if sess.IsDone() {
		err = h.loginSession.Delete(ctx, sess.Challenge)
		if err != nil {
			return nil, err.WithTrace("session.Delete")
		}
		return nil, nil
	}
	// store session
	err = h.loginSession.Save(ctx, sess)
	if err != nil {
		return nil, err.WithTrace("session.Store")
	}
	return r, nil
}
