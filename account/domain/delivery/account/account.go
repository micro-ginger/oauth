package account

import (
	"time"

	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
)

type Account[T account.ExtentedModel] struct {
	Id uint64 `json:"id"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`

	Status account.Status `json:"status"`

	Extended any `json:",inline"`
}

func NewAccount[T account.ExtentedModel](acc *a.Account[T]) *Account[T] {
	return &Account[T]{
		Id:        acc.GetId(),
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
		Status:    acc.GetStatus(),
		Extended:  acc.T.GetDeliveryResult(),
	}
}
