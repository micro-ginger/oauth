package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
)

func (uc *useCase[T]) Update(ctx context.Context,
	q query.Query, update *a.Update[T]) errors.Error {
	changed := false
	fq := query.NewUpdate(q)
	// internal status
	var internalStatuses account.InternalStatus = 0

	if update.UpdateStatus != nil {
		if update.UpdateStatus.AddStatus > 0 {
			fq.WithOr("status", update.UpdateStatus.AddStatus)
			changed = true
		}
		if update.UpdateStatus.DelStatus > 0 {
			fq.WithNot("status", update.UpdateStatus.DelStatus)
			changed = true
		}
		if update.UpdateStatus.AddStatus.Is(account.StatusRegistered) {
			internalStatuses |= account.InternalStatusRegistered
			changed = true
		}
	}

	if update.UpdateInternalStatus != nil {
		if update.UpdateInternalStatus.AddStatus > 0 {
			internalStatuses |= update.UpdateInternalStatus.AddStatus
			changed = true
		}
		if update.UpdateInternalStatus.DelStatus > 0 {
			fq.WithNot("internal_status", update.UpdateInternalStatus.DelStatus)
			changed = true
		}
	}
	if internalStatuses > 0 {
		fq.WithOr("internal_status", internalStatuses)
		changed = true
	}

	if update.UpdatePassword != nil {
		if err := uc.ValidatePassword(ctx, update.UpdatePassword.New); err != nil {
			return err
		}
		hashedPassword, err := a.HashPassword(update.UpdatePassword.New)
		if err != nil {
			return err
		}
		fq.WithSet("hashed_password", hashedPassword)
		changed = true
	}
	changed = update.T.ProcessUpdates(fq) || changed
	if changed {
		if err := uc.repo.Update(fq, nil); err != nil {
			return err.WithTrace("repo.Update")
		}
	}

	if update.Account != nil {
		if err := uc.UpdateAccount(ctx, q, update.Account); err != nil {
			return err.WithTrace("UpdateAccount")
		}
	}
	return nil
}

func (uc *useCase[T]) UpdateAccount(ctx context.Context,
	q query.Query, update *a.Account[T]) errors.Error {
	if err := uc.repo.Update(q, update); err != nil {
		return err.WithTrace("repo.Update")
	}
	return nil
}
