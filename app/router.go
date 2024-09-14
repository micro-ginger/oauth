package app

import (
	"github.com/ginger-core/gateway"
	"github.com/micro-blonde/auth/proto/auth"
	"github.com/micro-ginger/oauth/global"
)

func (a *App[acc, prof, regReq, reg, f]) registerRoutes() {
	a.registerHttpRoutes()
	a.registerGrpcRoutes()
}

func (a *App[acc, prof, regReq, reg, f]) registerHttpRoutes() {
	rg := a.HTTP.NewRouterGroup("/")
	//
	// chaptcha
	if a.Captcha != nil {
		rg.OnPath(
			gateway.Create,
			"/captcha/generate",
			a.Captcha.GenerateHandler,
		)
	}
	//
	// login
	loginGroup := rg.Group("/login")
	loginGroup.OnPath(
		gateway.Create,
		"",
		a.Authenticator.ShouldVerifyCaptcha(),
		a.Login.Handler,
	)
	//
	// accounts
	accountsGroup := rg.Group("/accounts")
	accountItemGroup := accountsGroup.Group("/:account_id")
	accountItemGroup.OnPath(
		gateway.Read,
		"",
		a.Authenticator.MustAuthenticate(),
		a.Account.GetHandler,
	)
	// account
	accountGroup := rg.Group("/account")
	accountGroup.OnPath(
		gateway.Read,
		"",
		a.Authenticator.MustHaveScope(global.ScopeReadAccount),
		a.Account.GetHandler,
	)
	accountGroup.OnPath(
		gateway.Update,
		"",
		a.Authenticator.MustHaveScope(global.ScopeUpdateProfile),
		a.Account.UpdateHandler,
	)
	// profile
	profileGroup := accountGroup.Group("/profile")
	profileGroup.OnPath(
		gateway.Read,
		"",
		a.Authenticator.MustHaveScope(global.ScopeReadProfile),
		a.Account.Profile.GetHandler,
	)
	profileGroup.OnPath(
		gateway.Update,
		"",
		a.Authenticator.MustHaveScope(global.ScopeUpdateProfile),
		a.Account.Profile.UpdateHandler,
	)
	profileGroup.OnPath(
		gateway.Create,
		"/photo",
		a.Authenticator.MustHaveScope(global.ScopeUpdateProfile),
		a.Account.Profile.PhotoUpdateHandler,
	)
	//
	// register
	registerGroup := rg.Group("/register")
	registerGroup.OnPath(
		gateway.Create,
		"",
		a.Authenticator.MustHaveScope(global.ScopeRegister),
		a.Register.RegisterHandler,
	)
	//
	// internal
	internalGroup := rg.Group("/internal")
	monitoringGroup := internalGroup.Group("/monitoring")
	monitoringGroup.OnPath(
		gateway.Read,
		"/health",
		a.Monitoring.Health,
	)
}

func (a *App[acc, prof, regReq, reg, f]) registerGrpcRoutes() {
	authGroup := a.GRPC.Register(&auth.Auth_ServiceDesc)
	authGroup.OnPath(
		gateway.Unknown,
		"ListAccounts",
		a.Account.GrpcListHandler,
	)
	authGroup.OnPath(
		gateway.Unknown,
		"GetAccount",
		a.Account.GrpcGetHandler,
	)
	authGroup.OnPath(
		gateway.Unknown,
		"ListProfiles",
		a.Account.Profile.GrpcListHandler,
	)
	authGroup.OnPath(
		gateway.Unknown,
		"GetProfile",
		a.Account.Profile.GrpcGetHandler,
	)
}
