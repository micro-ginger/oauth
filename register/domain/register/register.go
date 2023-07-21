package register

import "time"

type Model interface {
	GetDeliveryResult() any
}

type Register[T Model] struct {
	Id uint64

	CreatedAt time.Time
	UpdatedAt *time.Time

	AccountId uint64

	HashedPassword []byte

	T T `gorm:"embedded"`
}

func (t *Register[T]) TableName() string {
	return "registers"
}
