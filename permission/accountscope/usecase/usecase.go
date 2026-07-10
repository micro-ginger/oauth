package usecase

import (
	"github.com/ginger-core/log"
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/permission/accountscope/domain"
	"github.com/micro-ginger/oauth/permission/accountscope/domain/accountscope"
)

type useCase[SessionAccountDetail a.ExtentedModel] struct {
	logger log.Logger
	repo   domain.Repository

	refreshScopeHandlers []accountscope.CreatedScopeEventHandle
}

func New[SessionAccountDetail a.ExtentedModel](logger log.Logger,
	repo domain.Repository) domain.UseCase[SessionAccountDetail] {
	uc := &useCase[SessionAccountDetail]{
		logger:               logger,
		repo:                 repo,
		refreshScopeHandlers: make([]accountscope.CreatedScopeEventHandle, 0),
	}
	return uc
}

func (uc *useCase[SessionAccountDetail]) RegisterCreateEventHandle(
	h accountscope.CreatedScopeEventHandle,
) {
	uc.refreshScopeHandlers = append(uc.refreshScopeHandlers, h)
}
