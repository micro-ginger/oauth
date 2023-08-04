package domain

import (
	"github.com/micro-ginger/oauth/account/domain/account"
	ra "github.com/micro-ginger/oauth/register/domain/account"
	"github.com/micro-ginger/oauth/register/domain/register"
)

type UseCase[T register.Model, acc account.Model] interface {
	register.UseCase[T, acc]
	Initialize(account ra.UseCase[acc])
}
