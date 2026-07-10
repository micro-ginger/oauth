package usecase

import (
	"github.com/ginger-core/log"
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/register/domain"
	ra "github.com/micro-ginger/oauth/register/domain/account"
	"github.com/micro-ginger/oauth/register/domain/register"
)

type base[T register.Model, acc a.ExtentedModel] struct {
	logger log.Logger
	repo   register.Repository[T]

	account ra.UseCase[acc]
}

type useCase[T register.Model, acc a.ExtentedModel] struct {
	*base[T, acc]
}

type uc[T register.Model, acc a.ExtentedModel] struct {
	base *base[T, acc]
	register.UseCase[T, acc]
}

func New[T register.Model, acc a.ExtentedModel](logger log.Logger,
	repo register.Repository[T]) domain.UseCase[T, acc] {
	base := &base[T, acc]{
		logger: logger,
		repo:   repo,
	}
	uc := &uc[T, acc]{
		base: base,
		UseCase: &useCase[T, acc]{
			base: base,
		},
	}
	return uc
}

func (uc *uc[T, acc]) Initialize(account ra.UseCase[acc]) {
	uc.base.account = account
}

func (uc *useCase[T, acc]) Base() register.UseCase[T, acc] {
	return uc
}

func (uc *useCase[T, acc]) Wrap(w register.UseCase[T, acc]) {

}

func (uc *uc[T, acc]) Wrap(w register.UseCase[T, acc]) {
	uc.UseCase = w
}
