package account

import "time"

type Account[T any] struct {
	T  T `gorm:"embedded"`
	Id uint64

	CreatedAt time.Time
	UpdatedAt *time.Time

	Status Status

	HashedPassword []byte
}

func NewAccount[T any]() *Account[T] {
	return &Account[T]{}
}
