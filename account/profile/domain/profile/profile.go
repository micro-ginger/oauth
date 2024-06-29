package profile

import (
	"time"

	"github.com/micro-blonde/auth/profile"
)

type Profile[T profile.Model] struct {
	profile.Profile[T] `gorm:"embedded" json:",inline"`

	UpdatedAt *time.Time
}

func NewProfile[T profile.Model]() *Profile[T] {
	return new(Profile[T])
}

func (m *Profile[T]) TableName() string {
	return "profiles"
}

func (m *Profile[T]) GetDeliveryResult() any {
	return nil
}
