package delivery

import (
	"strings"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/login/domain/delivery/login"
	"github.com/micro-ginger/oauth/login/flow"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

const DefaultSection = "DEFAULT"

func (h *lh[acc]) getFlow(_ gateway.Request,
	req *login.Request) (*flow.Flow, errors.Error) {
	if req.Section == "" {
		req.Section = DefaultSection
	}
	flow := h.flows.Get(req.Section, req.Stage)
	if flow == nil {
		return nil, errors.Unauthorized().
			WithTrace("flows.Get.nil")
	}
	return flow, nil
}

func (h *lh[acc]) generateSession(request gateway.Request,
	flow *flow.Flow, req *login.Request) (*session.Session[acc], errors.Error) {
	stepQ, _ := request.GetQuery("step")
	genReq := &session.GenerateRequest{
		Flow:  flow,
		Step:  stepQ,
		Roles: strings.Split(req.Roles, ","),
	}
	return h.storeSession(request, genReq)
}

func (h *lh[acc]) storeSession(request gateway.Request,
	genReq *session.GenerateRequest) (*session.Session[acc], errors.Error) {
	session, err := h.loginSession.Generate(request.GetContext(), genReq)
	if err != nil {
		return nil, err.
			WithTrace("session.Generate")
	}
	return session, nil
}
