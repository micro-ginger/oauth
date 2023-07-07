package account

import (
	"time"

	"github.com/ginger-core/errors"
	"github.com/micro-blonde/auth/account"
	"golang.org/x/crypto/bcrypt"
)

type Model interface {
	account.Model
}

type Account[T Model] struct {
	account.Account[T]

	CreatedAt time.Time
	UpdatedAt *time.Time

	HashedPassword []byte
}

func NewAccount[T Model]() *Account[T] {
	return &Account[T]{}
}

func (m *Account[T]) GetDeliveryResult() any {
	return nil
}

func (a *Account[T]) MatchPassword(password string) error {
	if a.HashedPassword == nil {
		return errors.Validation().
			WithDesc("password is nil")
	}
	return bcrypt.CompareHashAndPassword(a.HashedPassword, []byte(password))
}
