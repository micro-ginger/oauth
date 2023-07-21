package info

import (
	"context"

	"github.com/ginger-core/errors"
)

func (h *handler[acc]) Delete(ctx context.Context, challenge string) errors.Error {
	return h.cache.Delete(ctx, h.getChallengeKey(challenge))
}

func (h *handler[acc]) DeleteItem(ctx context.Context, challenge, key string) errors.Error {
	return h.cache.UnsetItem(ctx, h.getChallengeKey(challenge), key)
}
