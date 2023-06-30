package usecase

import (
	"github.com/ginger-core/log"
	"github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/account/domain"
	a "github.com/micro-ginger/oauth/account/domain/account"
)

type useCase[T account.Model] struct {
	logger log.Logger
	repo   a.Repository[T]
}

func New[T account.Model](logger log.Logger,
	repo a.Repository[T]) domain.UseCase[T] {
	return &useCase[T]{
		logger: logger,
		repo:   repo,
	}
}
