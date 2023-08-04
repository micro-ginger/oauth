package register

import "github.com/micro-ginger/oauth/register/domain/register"

type Response[T register.Model] struct {
	Id uint64 `json:"id"`

	T any `json:"meta,inline"`
}

func NewResponse[T register.Model](reg *register.Register[T]) *Response[T] {
	return &Response[T]{
		Id: reg.Id,
		T:  reg.T.GetDeliveryResult(),
	}
}
