package repository

import (
	"github.com/ginger-core/repository"
	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
)

type repo[T account.Model] struct {
	repository.Repository
}

func New[T account.Model](base repository.Repository) a.Repository[T] {
	repo := &repo[T]{
		Repository: base,
	}
	return repo
}
