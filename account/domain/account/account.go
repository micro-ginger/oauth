package account

import (
	"time"

	"github.com/micro-blonde/auth/account"
)

type Account[T account.Model] struct {
	account.Account[T]

	CreatedAt time.Time
	UpdatedAt *time.Time

	HashedPassword []byte
}

func NewAccount[T account.Model]() *Account[T] {
	return &Account[T]{}
}
