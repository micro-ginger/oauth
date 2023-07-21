package usecase

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/account/domain"
	a "github.com/micro-ginger/oauth/account/domain/account"
)

type useCase[T account.Model] struct {
	logger log.Logger
	config config

	repo a.Repository[T]
}

func New[T account.Model](logger log.Logger, registry registry.Registry,
	repo a.Repository[T]) domain.UseCase[T] {
	uc := &useCase[T]{
		logger: logger,
		repo:   repo,
	}
	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.initialize()
	return uc
}
