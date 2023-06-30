package account

import (
	"time"

	"github.com/micro-blonde/auth/account"
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
