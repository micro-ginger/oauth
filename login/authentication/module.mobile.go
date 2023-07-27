package authentication

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log/logger"
	"github.com/micro-ginger/oauth/login/authentication/otp"
	"github.com/micro-ginger/oauth/login/authentication/step"
	sbase "github.com/micro-ginger/oauth/login/authentication/steps/base"
	mobileAcc "github.com/micro-ginger/oauth/login/authentication/steps/mobile/account"
	mobileOtp "github.com/micro-ginger/oauth/login/authentication/steps/mobile/otp"
	mobileVerify "github.com/micro-ginger/oauth/login/authentication/steps/mobile/verify"
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
	m.baseHandler = sbase.New(
		m.logger.WithTrace("base"),
		m.registry.ValueOf("base"),
		m.cache, m.info,
	)
	for _, cfg := range config.Steps {
		m.initializeHandler(
			m.registry.ValueOf(string(cfg.Type)), cfg.Type,
		)
	}
}

func (m *MobileModule[acc]) initializeHandler(
	registry registry.Registry, handlerType step.Type) {
	var h step.Handler[acc]
	switch handlerType {
	case mobileOtp.Type:
		otp := otp.New(
			m.logger.WithTrace("mobileOtp.otp"),
			registry.ValueOf(string(handlerType)),
			m.info,
			m.validator.Session, m.validator.Global,
		)
		h = mobileOtp.New(
			m.logger.WithTrace("mobileOtp"),
			registry.ValueOf(string(handlerType)),
			m.baseHandler, otp,
		)
	case mobileVerify.Type:
		otp := otp.New(
			m.logger.WithTrace("mobileVerify.otp"),
			registry.ValueOf(string(handlerType)),
			m.info,
			m.validator.Session, m.validator.Global,
		)
		h = mobileVerify.New(
			m.logger.WithTrace("mobileVerify"),
			registry.ValueOf(string(handlerType)),
			m.baseHandler, otp,
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
	m.steps.RegisterHandler(string(handlerType), h)
}
