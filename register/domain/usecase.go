package domain

import (
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/register/domain/register"
)

type UseCase[T register.Model, acc account.Model] interface {
	register.UseCase[T, acc]
}
