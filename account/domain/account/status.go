package account

type Status uint64

const (
	StatusVerified Status = 0b1 << iota
	StatusRejected
	StatusBlocked
	StatusDisabled
)

func (s Status) Uint64() uint64 {
	return uint64(s)
}

func (s *Status) Add(status Status) Status {
	*s |= status
	return *s
}

func (s Status) Is(status Status) bool {
	return s&status == status
}

func (s Status) Has(status Status) bool {
	return s&status > 0
}
