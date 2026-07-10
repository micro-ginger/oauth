package session

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/session/domain/session"
	r "github.com/micro-ginger/oauth/session/repository"
	"github.com/micro-ginger/oauth/session/usecase"
)

type Module[AccountDetail a.ExtentedModel] struct {
	Repository session.Repository[AccountDetail]
	UseCase    session.UseCase[AccountDetail]
}

func New[AccountDetail a.ExtentedModel](logger log.Logger, registry registry.Registry,
	cache repository.Cache) *Module[AccountDetail] {
	repoLogger := logger.WithTrace("repo")
	repo := r.New[AccountDetail](repoLogger, cache)

	ucLogger := logger.WithTrace("uc")
	uc := usecase.New(ucLogger, registry, repo)

	m := &Module[AccountDetail]{
		Repository: repo,
		UseCase:    uc,
	}
	return m
}
