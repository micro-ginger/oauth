package accountrole

import (
	"github.com/ginger-core/log"
	dl "github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/permission/accountrole/domain"
	"github.com/micro-ginger/oauth/permission/accountrole/repository"
	"github.com/micro-ginger/oauth/permission/accountrole/usecase"
)

type Module struct {
	Repository domain.Repository
	UseCase    domain.UseCase
}

func Initialize(logger log.Logger, baseDb dl.Repository) *Module {
	repo := repository.New(baseDb)
	uc := usecase.New(logger, repo)

	m := &Module{
		Repository: repo,
		UseCase:    uc,
	}
	return m
}
