package register

import (
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type Request[T Model, acc a.ExtentedModel] struct {
	Register *Register[T]
	Update   *account.Update[acc]
}
