package info

import (
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type Info[acc a.ExtentedModel] struct {
	AccountId uint64
	Account   *account.Account[acc] `json:"-"`

	RequestedRoles []string
	Section        string

	Temp Temp
}

func (i *Info[acc]) PopulateAccount(a *account.Account[acc]) {
	i.AccountId = a.GetId()
	i.Account = a
	// i.AccountStatus = a.Status.Uint64()
}

func NewFromAccount[acc a.ExtentedModel](a *account.Account[acc]) *Info[acc] {
	return &Info[acc]{
		AccountId: a.GetId(),
		Account:   a,
		// AccountStatus: a.Status.Uint64(),
	}
}

func New[acc a.ExtentedModel]() *Info[acc] {
	return &Info[acc]{
		Temp: make(Temp),
	}
}

func (i *Info[acc]) SetTemp(key string, value any) {
	if i.Temp == nil {
		i.Temp = make(map[string]any)
	}
	i.Temp.Set(key, value)
}

func (i *Info[acc]) GetTemp(key string) any {
	if i.Temp == nil {
		return nil
	}
	return i.Temp.Get(key)
}

func (i *Info[acc]) DelTemp(key string) {
	if i.Temp == nil {
		return
	}
	i.Temp.Del(key)
}
