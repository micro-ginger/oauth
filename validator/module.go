package validator

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	rd "github.com/micro-ginger/oauth/validator/domain/validator"
	r "github.com/micro-ginger/oauth/validator/repository"
	"github.com/micro-ginger/oauth/validator/usecase"
)

type Module struct {
	Repository rd.Repository
	UseCase    rd.UseCase
}

func New(logger log.Logger,
	registry registry.Registry, cache repository.Cache) *Module {
	repo := r.New(cache)
	uc := usecase.New(logger, registry, repo)
	m := &Module{
		Repository: repo,
		UseCase:    uc,
	}
	return m
}
