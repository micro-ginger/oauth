package session

import (
	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/login/authentication/step"
	"github.com/micro-ginger/oauth/login/flow"
)

type GenerateRequest struct {
	Flow *flow.Flow
	Step string
}

func (uc *useCase) Generate(request *GenerateRequest) (*Session, errors.Error) {
	session := &Session{
		Flow: Flow{
			Flow: request.Flow,
			Pos: Position{
				StageIndex:  0,
				StepIndex:   0,
				ActionIndex: 0,
			},
		},
	}

	if request.Step != "" {
		stepInd := session.Flow.Stages[session.Flow.Pos.StageIndex].
			GetStepIndex(step.Type(request.Step))
		if stepInd < 0 {
			return nil, errors.Validation().
				WithTrace("GetStepIndex.ind<0")
		}
		session.Flow.Pos.StepIndex = stepInd
	}
	return session, nil
}
