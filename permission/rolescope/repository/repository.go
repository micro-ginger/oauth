package repository

import (
	dl "github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/permission/rolescope/domain"
)

type repo struct {
	dl.Repository
}

func New(base dl.Repository) domain.Repository {
	repo := &repo{
		Repository: base,
	}
	return repo
}
