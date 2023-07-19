package repository

import (
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/register/domain/register"
)

type repo[T register.Model] struct {
	base repository.Transational
}

func New[T register.Model](base repository.Transational) register.Repository[T] {
	repo := &repo[T]{
		base: base,
	}
	return repo
}
