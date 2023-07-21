package info

import (
	"context"

	"github.com/ginger-core/errors"
)

func (h *handler[acc]) Generate(ctx context.Context) (*Info[acc], errors.Error) {
	key, err := h.challengeGenerator(h.config.Challenge.Characters, 10)
	if err != nil {
		return nil, err
	}

	inf := &Info[acc]{
		Key:        key,
		HandlerInd: 0,
	}
	if err := h.Save(ctx, inf); err != nil {
		return nil, err
	}
	if err := h.cache.Expire(ctx, h.getChallengeKey(inf.Challenge), h.config.InfoExpiration); err != nil {
		return nil, err
	}
	return inf, nil
}
