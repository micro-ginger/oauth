package domain

import (
	a "github.com/micro-blonde/auth/account"
	ra "github.com/micro-ginger/oauth/register/domain/account"
	"github.com/micro-ginger/oauth/register/domain/register"
)

type UseCase[T register.Model, acc a.ExtentedModel] interface {
	register.UseCase[T, acc]
	Initialize(account ra.UseCase[acc])
}
