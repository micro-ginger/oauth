package app

import (
	"github.com/ginger-repository/redis"
	"github.com/ginger-repository/sql"
)

func (a *app) initializeDatabases() {
	a.initializeSql()
	a.initializeRedis()
	a.initializeCache()
}

func (a *app) initializeSql() {
	sqlLogger := a.logger.WithTrace("sql")
	a.sql = sql.New(sqlLogger, a.registry.ValueOf("sql"))
	if err := a.sql.Initialize(); err != nil {
		panic(err)
	}
}

func (a *app) initializeRedis() {
	a.redis = redis.NewRepository(a.registry.ValueOf("redis"))
	if err := a.redis.Initialize(); err != nil {
		panic(err)
	}
}

func (a *app) initializeCache() {
	a.cache = redis.NewCache(a.redis)
}
