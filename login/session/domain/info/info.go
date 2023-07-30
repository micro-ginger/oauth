package info

import (
	"github.com/micro-ginger/oauth/account/domain/account"
)

type Info[acc account.Model] struct {
	AccountId uint64
	Account   *account.Account[acc] `json:"-"`

	RequestedRoles []string
	Section        string

	Temp map[string]any
}

func (i *Info[acc]) PopulateAccount(a *account.Account[acc]) {
	i.AccountId = a.Id
	i.Account = a
	// i.AccountStatus = a.Status.Uint64()
}

func NewFromAccount[acc account.Model](a *account.Account[acc]) *Info[acc] {
	return &Info[acc]{
		AccountId: a.Id,
		Account:   a,
		// AccountStatus: a.Status.Uint64(),
	}
}

func New[acc account.Model]() *Info[acc] {
	return &Info[acc]{}
}

func (i *Info[acc]) SetTemp(key string, value any) {
	if i.Temp == nil {
		i.Temp = make(map[string]any)
	}
	i.Temp[key] = value
}

func (i *Info[acc]) GetTemp(key string) any {
	if i.Temp == nil {
		return nil
	}
	return i.Temp[key]
}
