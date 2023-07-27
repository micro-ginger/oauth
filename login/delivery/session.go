package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/domain/delivery/login"
	"github.com/micro-ginger/oauth/login/session"
)

const DefaultSection = "DEFAULT"

func (h *lh[acc]) newSession(request gateway.Request,
	req *login.Request) (*session.Session, errors.Error) {
	if req.Section == "" {
		req.Section = DefaultSection
	}
	flow := h.flows.Get(req.Section)
	if flow == nil {
		return nil, errors.Unauthorized().
			WithTrace("flows.Get.nil")
	}

	stepQ, _ := request.GetQuery("step")

	genReq := &session.GenerateRequest{
		Flow: flow,
		Step: stepQ,
	}
	session, err := h.session.Generate(genReq)
	if err != nil {
		return nil, err.
			WithTrace("session.Generate")
	}

	return session, nil
}
