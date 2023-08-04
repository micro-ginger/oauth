package usecase

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/register/domain"
	a "github.com/micro-ginger/oauth/register/domain/account"
	"github.com/micro-ginger/oauth/register/domain/register"
)

type base[T register.Model, acc account.Model] struct {
	logger log.Logger
	repo   register.Repository[T]

	account a.UseCase[acc]
}

type useCase[T register.Model, acc account.Model] struct {
	*base[T, acc]
}

type uc[T register.Model, acc account.Model] struct {
	base *base[T, acc]
	register.UseCase[T, acc]
}

func New[T register.Model, acc account.Model](logger log.Logger,
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

func (uc *uc[T, acc]) Initialize(account a.UseCase[acc]) {
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
