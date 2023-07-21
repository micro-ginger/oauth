package repository

import (
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type repo struct {
	logger log.Logger
	cache  repository.Cache
}

func New(logger log.Logger, cache repository.Cache) session.Repository {
	repo := &repo{
		logger: logger,
		cache:  cache,
	}
	return repo
}
