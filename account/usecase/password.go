package usecase

import (
	"context"
	"unicode"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/account/domain/account"
	"golang.org/x/crypto/bcrypt"
)

func (uc *useCase[T]) ValidatePassword(
	ctx context.Context, password string) errors.Error {
	if len(password) < uc.config.Password.MinLen {
		return errors.Validation().
			WithContext(ctx).
			WithId("NewPasswordIsShortError").
			WithMessage("Password must contain at least {{.minLen}} characters.").
			WithProperty("minLen", uc.config.Password.MinLen)
	}
	var hasNumber, hasUpperChar, hasLowerChar, hasSpecial bool
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpperChar = true
		case unicode.IsLower(c):
			hasLowerChar = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}
	rate := 0
	if hasNumber {
		rate++
	}
	if hasUpperChar {
		rate++
	}
	if hasLowerChar {
		rate++
	}
	if hasSpecial {
		rate++
	}
	if rate < uc.config.Password.MinComplexity {
		return errors.Validation().
			WithContext(ctx).
			WithId("PasswordIsNotComplexError").
			WithMessage("Password is not complex enough.")
	}
	return nil
}

func (uc *useCase[T]) VerifyPassword(ctx context.Context,
	account *account.Account[T], password string) errors.Error {
	if len(account.HashedPassword) == 0 {
		return errors.Unauthorized().
			WithId("InvalidPasswordError").
			WithTrace("VerifyPassword.empty").
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
