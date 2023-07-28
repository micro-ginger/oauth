package authentication

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log/logger"
	"github.com/micro-ginger/oauth/login/authentication/otp"
	sbase "github.com/micro-ginger/oauth/login/authentication/steps/base"
	mobileAcc "github.com/micro-ginger/oauth/login/authentication/steps/mobile/account"
	mobileOtp "github.com/micro-ginger/oauth/login/authentication/steps/mobile/otp"
	mobileVerify "github.com/micro-ginger/oauth/login/authentication/steps/mobile/verify"
	"github.com/micro-ginger/oauth/login/flow/stage/step"
	"github.com/micro-ginger/oauth/login/flow/stage/step/handler"
	"github.com/micro-ginger/oauth/validator"
)

type MobileModule[acc mobileAcc.Model] struct {
	*Base[acc]
}

func NewMobile[acc mobileAcc.Model](base *Base[acc]) *MobileModule[acc] {
	m := &MobileModule[acc]{
		Base: base,
	}
	return m
}

func (m *MobileModule[acc]) Initialize() {
	m.Base.Initialize()
	m.initializeHandlers()
}

func (m *MobileModule[acc]) initializeHandlers() {
	config := new(config)
	if err := m.registry.Unmarshal(config); err != nil {
		panic(err)
	}
	for key, cfg := range config.Steps {
		m.initializeHandler(
			m.registry.ValueOf(key), cfg.Type,
		)
	}
}

func (m *MobileModule[acc]) initializeHandler(
	registry registry.Registry, handlerType step.Type) {
	baseHandler := sbase.New(
		m.logger.WithTrace("base"),
		m.registry.ValueOf("base"),
		m.loginSession,
		m.cache,
	)
	baseHandler.WithType(handlerType)

	sessionValidator := validator.New(
		m.logger.WithTrace("validators.session"),
		registry.ValueOf("validators.session"),
		m.cache,
	)
	globalValidator := validator.New(
		m.logger.WithTrace("validators.global"),
		registry.ValueOf("validators.global"),
		m.cache,
	)
	var h handler.Handler[acc]
	switch handlerType {
	case mobileOtp.Type:
		otp := otp.New(
			m.logger.WithTrace("mobileOtp.otp"),
			registry,
			m.loginSession,
			sessionValidator.UseCase,
			globalValidator.UseCase,
		)
		h = mobileOtp.New(
			m.logger.WithTrace("mobileOtp"),
			registry,
			baseHandler, otp,
		)
	case mobileVerify.Type:
		otp := otp.New(
			m.logger.WithTrace("mobileVerify.otp"),
			registry,
			m.loginSession,
			sessionValidator.UseCase,
			globalValidator.UseCase,
		)
		h = mobileVerify.New(
			m.logger.WithTrace("mobileVerify"),
			registry,
			baseHandler, otp,
		)
	default:
		m.logger.
			With(logger.Field{
				"type": handlerType,
			}).
			WithTrace("handler.notFound").
			Warnf("step handler not found")
		return
	}
	m.steps.RegisterHandler(handlerType, h)
}
