package account

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-blonde/auth/account"
	ad "github.com/micro-ginger/oauth/account/domain/account"
)

type UseCase[T account.Model] interface {
	GetById(ctx context.Context, id uint64) (*ad.Account[T], errors.Error)
}
