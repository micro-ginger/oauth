package handler

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/session/domain/flow"
	"github.com/micro-ginger/oauth/login/session/domain/info"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *handler[acc]) Generate(ctx context.Context,
	request *session.GenerateRequest) (*session.Session[acc], errors.Error) {
	key, err := h.challengeGenerator(h.config.Challenge.Characters, 10)
	if err != nil {
		return nil, err
	}

	sess := &session.Session[acc]{
		Key: key,
		Flow: flow.Flow{
			Flow: request.Flow,
			Pos: flow.Position{
				StageIndex:  0,
				StepIndex:   0,
				ActionIndex: 0,
			},
		},
		Info: info.New[acc](),
	}

	if request.Step != "" {
		stepInd := sess.Flow.Stages[sess.Flow.Pos.StageIndex].
			GetStepIndex(step.Type(request.Step))
		if stepInd < 0 {
			return nil, errors.Validation().
				WithTrace("GetStepIndex.ind<0")
		}
		sess.Flow.Pos.StepIndex = stepInd
	}
	return sess, nil
}
