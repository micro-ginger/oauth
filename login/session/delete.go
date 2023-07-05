package session

import (
	"context"

	"github.com/ginger-core/errors"
)

func (uc *useCase) Delete(ctx context.Context, session *Session) errors.Error {
	return uc.cache.Delete(ctx, session.GetKey())
}
