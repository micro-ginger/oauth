package otp

import (
	"context"
	"time"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/session/domain/session"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

type Handler interface {
	// RegisterCodeGenerator registers custom code generator for custom code generation
	RegisterCodeGenerator(generator func() string)

	Generate(ctx context.Context, key any, challenge string, otpType string) (*Otp, time.Duration, errors.Error)
	// Verify verifies the challenge with code.
	// metaRef is reference of meta existing in otp, which will be filled while loading the otp
	Verify(ctx context.Context, challenge string, otpType string, code string) (*Otp, errors.Error)
}

type handler[acc account.Model] struct {
	logger log.Logger
	config config

	session session.Handler[acc]

	codeGenerator func() string
	// sessionValidator is otp validation which is being validated in each login session
	sessionValidator validator.UseCase
	// globalValidator is validation for entity globally to handle validation of all actions of current entity type
	globalValidator validator.UseCase
}

func New[acc account.Model](logger log.Logger, registry registry.Registry,
	session session.Handler[acc], sessionValidator validator.UseCase,
	globalValidator validator.UseCase) Handler {
	h := &handler[acc]{
		logger:           logger,
		session:          session,
		sessionValidator: sessionValidator,
		globalValidator:  globalValidator,
	}
	if err := registry.Unmarshal(&h.config); err != nil {
		panic(err)
	}
	h.config.Initialize()

	h.RegisterCodeGenerator(h.GenerateCode)

	return h
}

func (h *handler[acc]) RegisterCodeGenerator(generator func() string) {
	h.codeGenerator = generator
}
