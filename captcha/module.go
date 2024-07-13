package captcha

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	cpt "github.com/micro-blonde/auth/captcha"
	"github.com/micro-ginger/oauth/captcha/delivery"
	"github.com/micro-ginger/oauth/captcha/domain/captcha"
	"github.com/micro-ginger/oauth/captcha/usecase"
)

type Module struct {
	UseCase         captcha.UseCase
	GenerateHandler gateway.Handler

	isInitialized bool
}

func New(logger log.Logger,
	registry registry.Registry, responder gateway.Responder) *Module {
	if registry == nil {
		return nil
	}
	ucLogger := logger.WithTrace("uc")
	uc := usecase.New(ucLogger, registry)

	m := &Module{
		UseCase: uc,
		GenerateHandler: delivery.NewGenerate(
			logger.WithTrace("delivery.generate"), uc, responder),
	}
	return m
}

func (m *Module) Initialize(generator cpt.Generator) {
	m.UseCase.Initialize(generator)
	m.isInitialized = true
}

func (m *Module) IsInitialized() bool {
	return m.isInitialized
}
