package key

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/domain/login"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
)

func (h *handler[T]) Process(request gateway.Request,
	meta *step.Meta) (any, errors.Error) {
	req := new(Request)
	if err := request.ProcessBody(req); err != nil {
		return nil, errors.
			Validation(err).
			WithTrace("request.ProcessBody")
	}

	ctx := request.GetContext()
	acc, err := h.getAccountHandleFunc(ctx, req.Key)
	if err != nil {
		return nil, err.
			WithTrace("getAccountHandleFunc")
	}

	if err := acc.MatchPassword(req.Password); err != nil {
		return nil, login.InvalidCredentialError.
			Clone().WithError(err).
			WithTrace("acc.MatchPassword")
	}

	return nil, nil
}
