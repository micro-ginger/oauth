package base

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
)

func (h *Handler[acc]) GetChallenge(ctx context.Context,
	request gateway.Request) (string, errors.Error) {
	ch, ok := request.GetQuery("challenge")
	if !ok {
		return ch, errors.Validation().
			WithTrace("GetChallenge.GetQuery").
			WithDesc("challenge is not passed")
	}
	return ch, nil
}
