package register

import "time"

type Model interface {
}

type Register[T Model] struct {
	Id uint64

	CreatedAt time.Time
	UpdatedAt *time.Time

	AccountId uint64

	HashedPassword []byte

	T T `gorm:"embedded" json:",inline"`
}

func (t *Register[T]) TableName() string {
	return "registers"
}
