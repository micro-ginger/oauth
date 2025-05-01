package repository

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type repo[AccountDetail gateway.ResultGetter] struct {
	logger log.Logger
	cache  repository.Cache
}

func New[AccountDetail gateway.ResultGetter](logger log.Logger,
	cache repository.Cache) session.Repository[AccountDetail] {
	repo := &repo[AccountDetail]{
		logger: logger,
		cache:  cache,
	}
	return repo
}
