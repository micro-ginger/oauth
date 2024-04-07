package role

import "time"

type State int

const (
	StateIsDefault State = 0b01
)

type Role struct {
	Id        uint64
	CreatedAt time.Time
	UpdatedAt *time.Time

	Name        string
	State       State
	Description string
}

func (m *Role) GetId() any {
	return m.Id
}
