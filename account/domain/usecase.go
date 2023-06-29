package domain

import "github.com/micro-ginger/oauth/account/domain/account"

type UseCase[T any] interface {
	account.UseCase[T]
}
