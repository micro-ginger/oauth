package session

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/session/domain/session"
	r "github.com/micro-ginger/oauth/session/repository"
	"github.com/micro-ginger/oauth/session/usecase"
)

type Module struct {
	Repository session.Repository
	UseCase    session.UseCase
}

func New(logger log.Logger, registry registry.Registry,
	cache repository.Cache) *Module {
	repoLogger := logger.WithTrace("repo")
	repo := r.New(repoLogger, cache)

	ucLogger := logger.WithTrace("uc")
	uc := usecase.New(ucLogger, registry, repo)

	m := &Module{
		Repository: repo,
		UseCase:    uc,
	}
	return m
}
