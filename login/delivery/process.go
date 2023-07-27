package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/info"
)

func (h *lh[acc]) process(request gateway.Request,
	info *info.Info[acc]) (any, errors.Error) {
	s, actionIdx := info.Session.Flow.GetCurrentStep()
	sh := h.stepHandlers[s.Type]
	if sh == nil {
		return nil, errors.Unauthorized().
			WithTrace("sh=nil").
			WithDesc("step handler not found")
	}

	r, err := sh.Process(request, info)
	if err != nil {
		return nil, err.
			WithDetail(errors.NewDetail().
				With("step", s).
				With("actionIdx", actionIdx)).
			WithTrace("sh.Process")
	}
	// next
	info.Session.Next()
	if info.Session.IsDone() {
		err = h.session.Delete(request.GetContext(), info.Session)
		if err != nil {
			return nil, err.WithTrace("session.Delete")
		}
		return nil, nil
	}
	// store session
	err = h.session.Store(request.GetContext(), info.Session)
	if err != nil {
		return nil, err.WithTrace("session.Store")
	}
	return r, nil
}
