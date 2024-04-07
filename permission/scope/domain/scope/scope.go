package scope

import "time"

type State int

const (
	StateIsDefault State = 0b01
)

type Scope struct {
	Id        uint64
	CreatedAt time.Time
	UpdatedAt *time.Time

	Name        string
	State       State
	Description string
}

func (m *Scope) GetId() any {
	return m.Id
}
