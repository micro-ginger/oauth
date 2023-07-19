package usecase

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/register/domain"
	a "github.com/micro-ginger/oauth/register/domain/account"
	"github.com/micro-ginger/oauth/register/domain/register"
)

type useCase[T register.Model, acc account.Model] struct {
	logger log.Logger
	repo   register.Repository[T]

	account a.UseCase[acc]
}

func New[T register.Model, acc account.Model](logger log.Logger,
	repo register.Repository[T]) domain.UseCase[T, acc] {
	uc := &useCase[T, acc]{
		logger: logger,
		repo:   repo,
	}
	return uc
}

func (uc *useCase[T, acc]) Initialize(account a.UseCase[acc]) {
	uc.account = account
}
