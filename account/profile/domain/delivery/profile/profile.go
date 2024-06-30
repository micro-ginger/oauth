package delivery

import (
	"github.com/micro-blonde/auth/profile"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type Profile[T profile.Model] struct {
	Id uint64 `json:"id"`
	T  any    `json:"detail"`
}

func NewProfile[T profile.Model](prof *p.Profile[T]) *Profile[T] {
	if prof == nil {
		return new(Profile[T])
	}
	return &Profile[T]{
		Id: prof.Id,
		T:  prof.T.GetDeliveryResult(),
	}
}

// func (h Profile[T]) MarshalJSON() ([]byte, error) {
// 	payloadJson, err := json.Marshal(h.T)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return payloadJson, nil
// }
