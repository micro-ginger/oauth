package usecase

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	cpt "github.com/micro-blonde/auth/captcha"
	"github.com/micro-ginger/oauth/captcha/domain/captcha"
)

type useCase struct {
	logger log.Logger
	config config

	generator cpt.Generator
}

func New(logger log.Logger, registry registry.Registry) captcha.UseCase {
	uc := &useCase{
		logger: logger,
	}
	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.initialize()
	return uc
}

func (uc *useCase) Initialize(generator cpt.Generator) {
	uc.generator = generator
}
