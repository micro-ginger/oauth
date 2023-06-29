package usecase

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type useCase[T any] struct {
	logger log.Logger
	repo   account.Repository[T]
}

func New[T any](logger log.Logger,
	repo account.Repository[T]) domain.UseCase[T] {
	return &useCase[T]{
		logger: logger,
		repo:   repo,
	}
}
