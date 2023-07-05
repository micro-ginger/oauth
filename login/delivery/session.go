package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	ldd "github.com/micro-ginger/oauth/login/domain/delivery/login"
	"github.com/micro-ginger/oauth/login/domain/login"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
)

func (h *lh) newSession(request gateway.Request,
	req *ldd.Request) (*login.Session, errors.Error) {

	flow := h.flows.Get(req.Section)
	if flow == nil {
		return nil, errors.Unauthorized().
			WithTrace("flows.Get.nil")
	}
	session := &login.Session{
		Flow: login.Flow{
			Flow: flow,
			Pos: login.Position{
				StageIndex:  0,
				StepIndex:   0,
				ActionIndex: 0,
			},
		},
	}

	stepQ, ok := request.GetQuery("step")
	if ok {
		stepInd := flow.Stages[session.Flow.Pos.StageIndex].
			GetStepIndex(step.Type(stepQ))
		if stepInd < 0 {
			return nil, errors.Validation().
				WithTrace("GetStepIndex.ind<0")
		}
		session.Flow.Pos.StepIndex = stepInd
	}
	return session, nil
}
