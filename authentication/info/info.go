package info

import (
	"encoding/json"

	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Info[acc account.Model] struct {
	Key any

	Challenge string `json:"-"`

	Account *account.Account[acc] `json:"-"`

	AccountId uint64

	Roles  []string
	Scopes []string

	Section       string
	AccountStatus uint64
	HandlerInd    int
	StepInd       int

	Session *session.Session

	Temp map[string]any
}

func (i Info[acc]) MarshalBinary() (data []byte, err error) {
	var bytes []byte
	bytes, err = json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (i *Info[acc]) PopulateAccount(a *account.Account[acc]) {
	i.Account = a
	i.AccountId = a.Id
	i.AccountStatus = a.Status.Uint64()
}

func NewFromAccount[acc account.Model](a *account.Account[acc]) *Info[acc] {
	return &Info[acc]{
		Account:       a,
		AccountId:     a.Id,
		AccountStatus: a.Status.Uint64(),
	}
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
