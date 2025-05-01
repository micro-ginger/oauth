package usecase

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/permission/accountscope/domain"
	"github.com/micro-ginger/oauth/permission/accountscope/domain/accountscope"
)

type useCase[SessionAccountDetail gateway.ResultGetter] struct {
	logger log.Logger
	repo   domain.Repository

	refreshScopeHandlers []accountscope.CreatedScopeEventHandle
}

func New[SessionAccountDetail gateway.ResultGetter](logger log.Logger,
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
