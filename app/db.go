package app

import (
	"github.com/ginger-repository/redis"
	"github.com/ginger-repository/sql"
)

func (a *App[acc, prof, regReq, reg, f]) initializeDatabases() {
	a.initializeSql()
	a.initializeRedis()
	a.initializeCache()
}

func (a *App[acc, prof, regReq, reg, f]) initializeSql() {
	sqlLogger := a.Logger.WithTrace("sql")
	a.Sql = sql.New(sqlLogger, a.Registry.ValueOf("sql"))
	if err := a.Sql.Initialize(); err != nil {
		panic(err)
	}
}

func (a *App[acc, prof, regReq, reg, f]) initializeRedis() {
	a.Redis = redis.NewRepository(a.Registry.ValueOf("redis"))
	if err := a.Redis.Initialize(); err != nil {
		panic(err)
	}
}

func (a *App[acc, prof, regReq, reg, f]) initializeCache() {
	a.Cache = redis.NewCache(a.Redis)
}
