package validator

import "time"

type Validation struct {
	Entries           []*Entry
	RequestedAt       time.Time
	RemainingVerifies int
	VerifyRequestedAt time.Time
	Key               any
}

func (v *Validation) Clone() *Validation {
	r := &Validation{
		Entries:           make([]*Entry, len(v.Entries)),
		RequestedAt:       v.RequestedAt,
		RemainingVerifies: v.RemainingVerifies,
	}
	for i, e := range v.Entries {
		t := e.Clone()
		r.Entries[i] = t
	}
	return r
}
