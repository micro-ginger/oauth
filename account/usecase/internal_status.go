package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
)

func (uc *useCase[T]) handleInternalStatus(ctx context.Context,
	acc *a.Account[T]) (err errors.Error) {
	var handledStatuses account.InternalStatus = 0
	defer func() {
		if handledStatuses > 0 {
			q := query.New(ctx)
			q = query.NewFilter(q).
				WithMatch(&query.Match{
					Key:      "id",
					Operator: query.Equal,
					Value:    acc.Id,
				})
			q = query.NewUpdate(q).
				WithNot("internal_status", handledStatuses)
			err = uc.repo.Update(q, nil)
		}
	}()
	for status, cfg := range uc.config.InternalStatus {
		if acc.InternalStatus.Has(status) {
			if len(cfg.AddRoles) > 0 {
				err := uc.accountRole.Assign(ctx, acc.Id, cfg.AddRoles)
				if err != nil {
					return err.WithTrace("accountRole.Assign")
				}
				handledStatuses |= status
			}
		}
	}
	if uc.manager != nil {
		oldStatus := acc.InternalStatus
		if err := uc.manager.HandleInternalStatus(ctx, acc); err != nil {
			return err.WithTrace("manager.HandleInternalStatus")
		}
		handledStatuses = oldStatus ^ acc.InternalStatus
	}
	return nil
}
