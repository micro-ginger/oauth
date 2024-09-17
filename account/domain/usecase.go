package domain

import (
	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/account/domain/permission"
)

type UseCase[T account.Model] interface {
	a.UseCase[T]
	Initialize(accountRole permission.AccountRole)
	SetManager(manager a.Manager[T])
	Start()
	Stop()
}
