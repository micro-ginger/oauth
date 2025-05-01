package accountscope

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	dl "github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/permission/accountscope/domain"
	"github.com/micro-ginger/oauth/permission/accountscope/repository"
	"github.com/micro-ginger/oauth/permission/accountscope/usecase"
)

type Module[SessionAccountDetail gateway.ResultGetter] struct {
	Repository domain.Repository
	UseCase    domain.UseCase[SessionAccountDetail]
}

func Initialize[SessionAccountDetail gateway.ResultGetter](logger log.Logger, baseDb dl.Repository) *Module[SessionAccountDetail] {
	repo := repository.New(baseDb)
	uc := usecase.New[SessionAccountDetail](logger, repo)

	m := &Module[SessionAccountDetail]{
		Repository: repo,
		UseCase:    uc,
	}
	return m
}
