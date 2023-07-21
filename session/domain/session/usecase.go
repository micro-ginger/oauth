package session

import (
	"context"

	"github.com/ginger-core/errors"
)

type UseCase interface {
	RegisterSessionHandlers(handlerFuncs ...SessionHandlerFunc)

	Create(ctx context.Context, session *CreateRequest) (*Session, errors.Error)

	ListSessionIdsOfAccount(ctx context.Context, accId uint64) ([]string, errors.Error)
	Get(ctx context.Context, accId uint64, id string) (*Session, errors.Error)
	GetByAccess(ctx context.Context, token string) (*Session, errors.Error)
	GetByRefresh(ctx context.Context, token string) (*Session, errors.Error)

	DeleteAll(ctx context.Context, session *Session) errors.Error
	Delete(ctx context.Context, accId uint64, id string) errors.Error
	DeleteAccess(ctx context.Context, token string) errors.Error
	DeleteRefresh(ctx context.Context, token string) errors.Error
}
