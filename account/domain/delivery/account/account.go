package account

import (
	"time"

	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
)

type Account[T account.Model] struct {
	Id uint64 `json:"id"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`

	Status account.Status `json:"status"`

	HashedPassword []byte `json:"-"`

	T any `gorm:"embedded"`
}

func NewAccount[T account.Model](acc *a.Account[T]) *Account[T] {
	r := &Account[T]{
		Id:        acc.Id,
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
		Status:    acc.Status,
	}
	r.T = acc.T.GetDeliveryResult()
	return r
}
