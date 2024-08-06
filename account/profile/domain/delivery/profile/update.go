package delivery

import (
	"github.com/micro-blonde/auth/profile"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type Update[T profile.Model] struct {
	T T `json:"detail"`
}

func (m *Update[T]) GetProfile() *p.Profile[T] {
	return &p.Profile[T]{
		Profile: profile.Profile[T]{
			T: m.T,
		},
	}
}
