package account

import (
	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
)

type Update[T a.Model] struct {
	T T `json:"detail"`
}

func (m *Update[T]) GetUpdate() *a.Update[T] {
	return &a.Update[T]{
		Account: &a.Account[T]{
			Account: account.Account[T]{
				T: m.T,
			},
		},
	}
}
