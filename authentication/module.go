package authentication

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/ginger-repository/redis"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/authentication/handler"
	"github.com/micro-ginger/oauth/authentication/info"
	"github.com/micro-ginger/oauth/login/session"
)

type Module[acc account.Model] struct {
	Info    info.Handler[acc]
	Handler *handler.Module[acc]
}

func Initialize[acc account.Model](logger log.Logger, registry registry.Registry,
	account account.UseCase[acc], session session.UseCase) *Module[acc] {
	m := &Module[acc]{}

	redisDb := redis.NewRepository(registry.ValueOf("redis"))
	if err := redisDb.Initialize(); err != nil {
		panic(err)
	}
	cache := redis.NewCache(redisDb)

	m.Info = info.New[acc](
		logger.WithTrace("info"),
		registry.ValueOf("info"),
		cache,
	)

	m.Handler = handler.New[acc](logger.WithTrace("handlers"))
	return m
}
