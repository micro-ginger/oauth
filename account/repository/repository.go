package repository

import (
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type repo[T any] struct {
	repository.Repository
}

func New[T any](base repository.Repository) account.Repository[T] {
	repo := &repo[T]{
		Repository: base,
	}
	return repo
}
