package session

type State int8

const (
	StateNone State = iota << 1
	StateFromDB
)

func (s *State) Add(s2 State) {
	*s = *s | s2
}

func (s State) Is(s2 State) bool {
	return s&s2 == s2
}

func (s State) Has(s2 State) bool {
	return s&s2 > 0
}
