package account

import a "github.com/micro-blonde/auth/account"

type Model interface {
	a.ExtentedModel
	GetMobile() *string
	MaskMobile() string
}
