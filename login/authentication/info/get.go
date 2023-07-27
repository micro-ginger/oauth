package info

import (
	"context"

	"github.com/ginger-core/errors"
)

func (h *handler[acc]) GetItem(ctx context.Context,
	challenge, key string, ref any) errors.Error {
	if err := h.cache.GetItem(ctx,
		h.getChallengeKey(challenge), key, ref); err != nil {
		return err
	}
	return nil
}

func (h *handler[acc]) Get(ctx context.Context,
	challenge string) (*Info[acc], errors.Error) {
	inf := new(Info[acc])
	if err := h.GetItem(ctx, challenge, root, inf); err != nil {
		return nil, err
	}
	inf.Challenge = challenge
	return inf, nil
}
