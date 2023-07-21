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
	fq := query.NewUpdate(q)
	// internal status
	var internalStatuses account.InternalStatus = 0

	if update.UpdateStatus != nil {
		if update.UpdateStatus.AddStatus > 0 {
			fq.WithOr("status", update.UpdateStatus.AddStatus)
		}
		if update.UpdateStatus.DelStatus > 0 {
			fq.WithNot("status", update.UpdateStatus.DelStatus)
		}
		if update.UpdateStatus.AddStatus.Is(account.StatusRegistered) {
			// registered
			internalStatuses |= account.InternalStatusRegistered
		}
		if update.UpdateStatus.AddStatus.Is(account.StatusVerified) {
			// verified
			internalStatuses |= account.InternalStatusVerified
		}
		if update.UpdateStatus.AddStatus.Is(account.StatusRejected) {
			// rejected
			internalStatuses |= account.InternalStatusRejected
		}
	}

	if update.UpdateInternalStatus != nil {
		if update.UpdateInternalStatus.AddStatus > 0 {
			internalStatuses |= update.UpdateInternalStatus.AddStatus
		}
		if update.UpdateInternalStatus.DelStatus > 0 {
			fq.WithNot("internal_status", update.UpdateInternalStatus.DelStatus)
		}
	}
	if internalStatuses > 0 {
		fq.WithOr("internal_status", internalStatuses)
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
	}

	if err := uc.repo.Update(q, nil); err != nil {
		return err.WithTrace("repo.Update")
	}
	// if update.UpdatePassword != nil {
	// 	// changed password
	// }

	if update.Account != nil {
		if err := uc.UpdateAccount(ctx, q, update.Account); err != nil {
			return err.WithTrace("UpdateAccount")
		}
	}
	return nil
	return nil
}

func (uc *useCase[T]) UpdateAccount(ctx context.Context,
	q query.Query, update *a.Account[T]) errors.Error {
	if err := uc.repo.Update(q, update); err != nil {
		return err.WithTrace("repo.Update")
	}
	return nil
}
