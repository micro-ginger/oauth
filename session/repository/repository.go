package repository

import (
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type repo[AccountDetail a.ExtentedModel] struct {
	logger log.Logger
	cache  repository.Cache
}

func New[AccountDetail a.ExtentedModel](logger log.Logger,
	cache repository.Cache) session.Repository[AccountDetail] {
	repo := &repo[AccountDetail]{
		logger: logger,
		cache:  cache,
	}
	return repo
}
