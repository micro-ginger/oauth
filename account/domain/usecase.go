package domain

import (
	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
)

type UseCase[T account.Model] interface {
	a.UseCase[T]
}
