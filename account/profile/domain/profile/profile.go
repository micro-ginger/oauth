package profile

import (
	"database/sql/driver"
	"encoding/json"
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

func (m Profile[T]) Value() (driver.Value, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (m *Profile[T]) Scan(src any) error {
	if src == nil {
		return nil
	}
	data, ok := src.([]byte)
	if !ok {
		return nil
	}
	d := new(Profile[T])
	err := json.Unmarshal(data, d)
	if err != nil {
		return err
	}
	*m = *d
	return nil
}

func (m *Profile[T]) TableName() string {
	return "profiles"
}

func (m *Profile[T]) GetDeliveryResult() any {
	return nil
}
