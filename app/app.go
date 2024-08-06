package app

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	redisRepo "github.com/ginger-repository/redis/repository"
	"github.com/ginger-repository/sql"
	"github.com/micro-blonde/auth/authorization"
	"github.com/micro-blonde/auth/profile"
	"github.com/micro-blonde/file"
	fileClient "github.com/micro-blonde/file/client"
	a "github.com/micro-ginger/oauth/account"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/captcha"
	"github.com/micro-ginger/oauth/login"
	"github.com/micro-ginger/oauth/permission"
	r "github.com/micro-ginger/oauth/register"
	rdd "github.com/micro-ginger/oauth/register/domain/delivery/register"
	"github.com/micro-ginger/oauth/register/domain/register"
	"github.com/micro-ginger/oauth/session"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Application interface {
	Initialize()
	Start()
}

type App[acc account.Model, prof profile.Model,
	regReq rdd.RequestModel, reg register.Model, f file.Model] struct {
	Registry registry.Registry
	Config   config
	Logger   log.Handler
	Language *i18n.Bundle
	/* database */
	Sql   sql.Repository
	Redis redisRepo.Repository
	Cache repository.Cache
	/* services */
	File fileClient.Client[f]
	/* modules */
	Captcha    *captcha.Module
	Account    *a.Module[acc, prof, f]
	permission *permission.Module
	Session    *session.Module
	Login      *login.Module[acc]
	Register   *r.Module[regReq, reg, acc]
	/* server */
	Authenticator authorization.Authenticator[acc]
	Ginger        gateway.Server
	GRPC          GrpcServer
}

func New[acc account.Model, prof profile.Model,
	regReq rdd.RequestModel, reg register.Model, f file.Model](
	configType string) *App[acc, prof, regReq, reg, f] {
	a := &App[acc, prof, regReq, reg, f]{
		Language: i18n.NewBundle(language.English),
	}
	a.loadConfig(configType)

	if err := a.Registry.Unmarshal(&a.Config); err != nil {
		panic(err)
	}
	return a
}

func (a *App[acc, prof, regReq, reg, f]) Initialize() {
	a.initializeLogger()
	a.initializeLanguage()
	a.initializeServer()
	a.initializeServices()
	a.initializeDatabases()
	a.initializeModules()
	a.initializeGrpc()
	a.registerRoutes()
}
