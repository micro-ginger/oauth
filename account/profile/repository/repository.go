package repository

import (
	"github.com/ginger-core/repository"
	prof "github.com/micro-blonde/auth/profile"
	"github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type repo[T prof.Model] struct {
	repository.Repository
}

func New[T prof.Model](base repository.Repository) profile.Repository[T] {
	repo := &repo[T]{
		Repository: base,
	}
	return repo
}
