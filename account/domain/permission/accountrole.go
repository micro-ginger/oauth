package permission

import (
	"context"

	"github.com/ginger-core/errors"
)

type AccountRole interface {
	Assign(ctx context.Context, accId uint64, roles []string) errors.Error
}
