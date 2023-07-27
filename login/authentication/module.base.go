package authentication

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-core/repository"
	"github.com/ginger-repository/redis"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/authentication/info"
	"github.com/micro-ginger/oauth/login/authentication/step"
	"github.com/micro-ginger/oauth/login/authentication/steps"
	sbase "github.com/micro-ginger/oauth/login/authentication/steps/base"
	keyPw "github.com/micro-ginger/oauth/login/authentication/steps/key/password"
	"github.com/micro-ginger/oauth/login/authentication/steps/password"
	"github.com/micro-ginger/oauth/login/authentication/steps/refresh"
	"github.com/micro-ginger/oauth/session/domain/session"
	"github.com/micro-ginger/oauth/validator/domain/validator"
)

type Base[acc account.Model] struct {
	logger   log.Logger
	registry registry.Registry

	cache repository.Cache

	info info.Handler[acc]

	validator struct {
		Session validator.UseCase
		Global  validator.UseCase
	}

	baseHandler *sbase.Handler[acc]
	steps       *steps.Module[acc]

	session session.UseCase
}

func newBase[acc account.Model](logger log.Logger, registry registry.Registry,
	account account.UseCase[acc], session session.UseCase) *Base[acc] {
	m := &Base[acc]{
		logger:   logger,
		registry: registry,
		session:  session,
	}

	redisDb := redis.NewRepository(registry.ValueOf("redis"))
	if err := redisDb.Initialize(); err != nil {
		panic(err)
	}
	m.cache = redis.NewCache(redisDb)
	return m
}

func (m *Base[acc]) Initialize() {
	m.info = info.New[acc](
		m.logger.WithTrace("info"),
		m.registry.ValueOf("info"),
		m.cache,
	)

	m.steps = steps.New[acc](m.logger.WithTrace("handlers"))

	m.initializeHandlers()
}

func (m *Base[acc]) initializeHandlers() {
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

func (m *Base[acc]) initializeHandler(
	registry registry.Registry, handlerType step.Type) {
	var h step.Handler[acc]
	switch handlerType {
	case keyPw.Type:
		h = keyPw.New(
			m.logger.WithTrace("keyPassword"),
			registry.ValueOf(string(handlerType)),
			m.baseHandler, m.cache,
		)
	case password.Type:
		h = password.New(
			m.logger.WithTrace("password"),
			registry.ValueOf(string(handlerType)),
			m.baseHandler,
		)
	case refresh.Type:
		h = refresh.New(
			m.logger.WithTrace("refresh"),
			m.baseHandler, m.session,
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
