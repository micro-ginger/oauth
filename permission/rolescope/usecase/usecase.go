package usecase

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/permission/rolescope/domain"
)

type useCase struct {
	logger log.Logger
	repo   domain.Repository
}

func New(logger log.Logger, repo domain.Repository) domain.UseCase {
	uc := &useCase{
		logger: logger,
		repo:   repo,
	}
	return uc
}
