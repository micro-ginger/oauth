package register

import (
	"github.com/micro-ginger/oauth/register/domain/register"
)

type Register[T register.Model] struct {
	Id uint64 `json:"id"`

	T any `json:",inline"`
}

func NewRegister[T register.Model](reg *register.Register[T]) *Register[T] {
	return &Register[T]{
		Id: reg.Id,
		T:  reg.T.GetDeliveryResult(),
	}
}
