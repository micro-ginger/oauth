package repository

import (
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

type repo struct {
	cache repository.Cache
}

func New(cache repository.Cache) validator.Repository {
	repo := &repo{
		cache: cache,
	}
	return repo
}
