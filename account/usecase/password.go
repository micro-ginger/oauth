package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/account/domain/account"
	"golang.org/x/crypto/bcrypt"
)

func (uc *useCase[T]) VerifyPassword(ctx context.Context,
	account *account.Account[T], password string) errors.Error {
	if len(account.HashedPassword) == 0 {
		return errors.Unauthorized().
			WithId("InvalidPasswordError").
			WithTrace("VerifyPassword.0").
			WithDesc("empty password").
			WithMessage("You have entered an invalid password.")
	}
	if err := bcrypt.CompareHashAndPassword(
		account.HashedPassword, []byte(password)); err != nil {
		return errors.Unauthorized(err).
			WithId("InvalidPasswordError").
			WithTrace("bcrypt.CompareHashAndPassword.Err").
			WithDesc("invalid password").
			WithMessage("You have entered an invalid password.")
	}
	return nil
}
