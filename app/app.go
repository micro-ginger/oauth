package app

import (
	"math/rand"
	"time"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	redisRepo "github.com/ginger-repository/redis/repository"
	"github.com/ginger-repository/sql"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Application interface {
	Initialize()
	Start()
}

type app struct {
	registry registry.Registry
	config   config
	logger   log.Handler
	language *i18n.Bundle
	/* database */
	sql   sql.Repository
	redis redisRepo.Repository
	cache repository.Cache
	/* services */
	/* modules */
	/* server */
	ginger gateway.Server
}

func New(configType string) Application {
	a := &app{
		language: i18n.NewBundle(language.English),
	}
	a.loadConfig(configType)

	if err := a.registry.Unmarshal(&a.config); err != nil {
		panic(err)
	}
	return a
}

func (a *app) Initialize() {
	rand.Seed(time.Now().UnixNano())

	a.initializeLogger()
	a.initializeLanguage()
	a.initializeServer()
	a.initializeServices()
	a.initializeDatabases()
	a.initializeModules()
	a.registerRoutes()
}
