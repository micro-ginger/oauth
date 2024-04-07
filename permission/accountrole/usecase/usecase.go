package usecase

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/permission/accountrole/domain"
	"github.com/micro-ginger/oauth/permission/accountrole/domain/accountrole"
)

type useCase struct {
	logger log.Logger
	repo   domain.Repository

	createRoleHandlers []accountrole.CreatedRoleEventHandle
}

func New(logger log.Logger, repo domain.Repository) domain.UseCase {
	uc := &useCase{
		logger:             logger,
		repo:               repo,
		createRoleHandlers: make([]accountrole.CreatedRoleEventHandle, 0),
	}
	return uc
}

func (uc *useCase) RegisterCreateEventHandle(h accountrole.CreatedRoleEventHandle) {
	uc.createRoleHandlers = append(uc.createRoleHandlers, h)
}
