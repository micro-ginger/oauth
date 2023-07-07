package app

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	redisRepo "github.com/ginger-repository/redis/repository"
	"github.com/ginger-repository/sql"
	"github.com/micro-blonde/auth/authorization"
	a "github.com/micro-ginger/oauth/account"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Application interface {
	Initialize()
	Start()
}

type App[acc account.Model] struct {
	Registry registry.Registry
	Config   config
	Logger   log.Handler
	Language *i18n.Bundle
	/* database */
	Sql   sql.Repository
	Redis redisRepo.Repository
	Cache repository.Cache
	/* services */
	/* modules */
	Account *a.Module[acc]
	Login   *login.Module
	/* server */
	Authenticator authorization.Authenticator[acc]
	Ginger        gateway.Server
}

func New[acc account.Model](configType string) *App[acc] {
	a := &App[acc]{
		Language: i18n.NewBundle(language.English),
	}
	a.loadConfig(configType)

	if err := a.Registry.Unmarshal(&a.Config); err != nil {
		panic(err)
	}
	return a
}

func (a *App[acc]) Initialize() {
	a.initializeLogger()
	a.initializeLanguage()
	a.initializeServer()
	a.initializeServices()
	a.initializeDatabases()
	a.initializeModules()
	a.registerRoutes()
}
