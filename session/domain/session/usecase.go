package session

import (
	"context"

	"github.com/ginger-core/errors"
	a "github.com/micro-blonde/auth/account"
)

type UseCase[AccountDetail a.ExtentedModel] interface {
	RegisterSessionHandlers(handlerFuncs ...SessionHandlerFunc[AccountDetail])

	Create(
		ctx context.Context, session *CreateRequest[AccountDetail],
	) (*Session[AccountDetail], errors.Error)

	ListSessionIdsOfAccount(ctx context.Context, accId uint64) ([]string, errors.Error)
	Get(ctx context.Context, accId uint64, id string) (*Session[AccountDetail], errors.Error)
	GetByAccess(ctx context.Context, token string) (*Session[AccountDetail], errors.Error)
	GetByRefresh(ctx context.Context, token string) (*Session[AccountDetail], errors.Error)

	DeleteAll(ctx context.Context, session *Session[AccountDetail]) errors.Error
	Delete(ctx context.Context, accId uint64, id string) errors.Error
	DeleteAccess(ctx context.Context, token string) errors.Error
	DeleteRefresh(ctx context.Context, token string) errors.Error
}
