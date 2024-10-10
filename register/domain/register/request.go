package register

import "github.com/micro-ginger/oauth/account/domain/account"

type Request[T Model, acc account.Model] struct {
	Register *Register[T]
	Update   *account.Update[acc]
}
