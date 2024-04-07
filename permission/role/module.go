package role

import (
	"github.com/ginger-core/log"
	dl "github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/permission/role/domain"
	"github.com/micro-ginger/oauth/permission/role/repository"
	"github.com/micro-ginger/oauth/permission/role/usecase"
)

type Module struct {
	Repository domain.Repository
	UseCase    domain.UseCase
}

func Initialize(logger log.Logger, baseDb dl.Repository) *Module {
	repo := repository.New(baseDb)
	uc := usecase.New(logger.WithTrace("uc"), repo)
	m := &Module{
		Repository: repo,
		UseCase:    uc,
	}
	return m
}
