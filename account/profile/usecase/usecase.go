package usecase

import (
	"github.com/ginger-core/log"
	prof "github.com/micro-blonde/auth/profile"
	"github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type useCase[T prof.Model] struct {
	logger log.Logger

	repo profile.Repository[T]
}

func New[T prof.Model](logger log.Logger,
	repo profile.Repository[T]) profile.UseCase[T] {
	uc := &useCase[T]{
		logger: logger,
		repo:   repo,
	}
	return uc
}
