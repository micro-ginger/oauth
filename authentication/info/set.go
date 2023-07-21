package info

import (
	"context"

	"github.com/ginger-core/errors"
)

func (h *handler[acc]) Set(ctx context.Context,
	challenge, key string, value any) errors.Error {
	if err := h.cache.SetItem(ctx,
		h.getChallengeKey(challenge), key, value); err != nil {
		return err
	}
	return nil
}

func (h *handler[acc]) Save(ctx context.Context, info *Info[acc]) errors.Error {
	challenge, err := h.challengeGenerator(
		h.config.Challenge.Characters,
		h.config.Challenge.Length,
	)
	if err != nil {
		return err
	}

	if info.Challenge != "" {
		err := h.cache.Rename(ctx,
			h.getChallengeKey(info.Challenge),
			h.getChallengeKey(challenge))
		if err != nil {
			return err
		}
	}

	info.Challenge = challenge

	return h.Set(ctx, challenge, root, info)
}
