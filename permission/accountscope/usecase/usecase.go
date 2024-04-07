package usecase

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/permission/accountscope/domain"
	"github.com/micro-ginger/oauth/permission/accountscope/domain/accountscope"
)

type useCase struct {
	logger log.Logger
	repo   domain.Repository

	refreshScopeHandlers []accountscope.CreatedScopeEventHandle
}

func New(logger log.Logger, repo domain.Repository) domain.UseCase {
	uc := &useCase{
		logger:               logger,
		repo:                 repo,
		refreshScopeHandlers: make([]accountscope.CreatedScopeEventHandle, 0),
	}
	return uc
}

func (uc *useCase) RegisterCreateEventHandle(h accountscope.CreatedScopeEventHandle) {
	uc.refreshScopeHandlers = append(uc.refreshScopeHandlers, h)
}
