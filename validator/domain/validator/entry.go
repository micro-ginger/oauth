package validator

import "time"

type Entry struct {
	MaxRequests       int
	MaxVerifies       int
	RemainingRequests int
	RemainingVerifies int
	ExpirationTime    time.Time
}

func (e *Entry) Clone() *Entry {
	return &Entry{
		MaxRequests:       e.MaxRequests,
		MaxVerifies:       e.MaxVerifies,
		RemainingRequests: e.RemainingRequests,
		RemainingVerifies: e.RemainingVerifies,
	}
}
