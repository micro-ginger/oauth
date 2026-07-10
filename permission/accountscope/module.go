package accountscope

import (
	"github.com/ginger-core/log"
	dl "github.com/ginger-core/repository"
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/permission/accountscope/domain"
	"github.com/micro-ginger/oauth/permission/accountscope/repository"
	"github.com/micro-ginger/oauth/permission/accountscope/usecase"
)

type Module[SessionAccountDetail a.ExtentedModel] struct {
	Repository domain.Repository
	UseCase    domain.UseCase[SessionAccountDetail]
}

func Initialize[SessionAccountDetail a.ExtentedModel](logger log.Logger, baseDb dl.Repository) *Module[SessionAccountDetail] {
	repo := repository.New(baseDb)
	uc := usecase.New[SessionAccountDetail](logger, repo)

	m := &Module[SessionAccountDetail]{
		Repository: repo,
		UseCase:    uc,
	}
	return m
}
