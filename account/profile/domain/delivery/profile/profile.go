package delivery

import (
	"github.com/micro-blonde/auth/profile"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type Profile[T profile.Model] struct {
	Id    uint64 `json:"id"`
	Photo string `json:"photo"`
	T     any    `json:"detail"`
}

func NewProfile[T profile.Model](prof *p.Profile[T]) *Profile[T] {
	if prof == nil {
		return new(Profile[T])
	}
	return &Profile[T]{
		Id:    prof.Id,
		Photo: prof.Photo,
		T:     prof.T.GetDeliveryResult(),
	}
}

type ProfilePhoto[T profile.Model] struct {
	Photo string `json:"photo"`
}

func NewProfilePhoto[T profile.Model](prof *p.Profile[T]) *ProfilePhoto[T] {
	if prof == nil {
		return new(ProfilePhoto[T])
	}
	return &ProfilePhoto[T]{
		Photo: prof.Photo,
	}
}
