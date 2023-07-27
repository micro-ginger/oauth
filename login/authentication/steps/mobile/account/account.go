package account

import "github.com/micro-ginger/oauth/account/domain/account"

type Model interface {
	account.Model
	GetMobile() *string
	MaskMobile() string
}
