package verify

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/authentication/response"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

type verifyBody struct {
	Code         string `json:"code" binding:"required"`
	NationalCode string `json:"nationalCode"`
}

func (h *_handler[acc]) Process(request gateway.Request,
	sess *session.Session[acc]) (response.Response, errors.Error) {
	ctx := request.GetContext()
	if sess.Flow.Pos.StepIndex == 0 {
		return h.request(ctx, request, sess)
	}
	r, err := h.verify(ctx, request, sess)
	if err != nil {
		return nil, err
	}
	return r, nil
}
