package app

import (
	"github.com/micro-ginger/oauth/account"
	"github.com/micro-ginger/oauth/captcha"
	"github.com/micro-ginger/oauth/login"
	"github.com/micro-ginger/oauth/monitoring"
	"github.com/micro-ginger/oauth/permission"
	"github.com/micro-ginger/oauth/register"
	"github.com/micro-ginger/oauth/session"
)

func (a *App[acc, prof, regReq, reg, f]) initializeModules() {
	a.initiateCaptcha()
	a.initiateAccount()
	a.initializePermission()
	a.initiateSession()
	a.initiateLogin()
	a.initiateRegister()
	a.initializeMonitoring()
	//
	// session create handlers
	a.Session.UseCase.RegisterSessionHandlers(
		a.Permission.AccountScope.UseCase.SessionAddRequestedRoleScopes,
		a.Permission.AccountScope.UseCase.SessionRemoveUnauthorized)
	//
	a.Register.Initialize(a.Account.UseCase)
	a.Account.Initialize(a.File)
	//
	a.Monitoring.Initialize(a.Redis, a.Sql)
}

func (a *App[acc, prof, regReq, reg, f]) initiateCaptcha() {
	a.Captcha = captcha.New(
		a.Logger.WithTrace("captcha"),
		a.Registry.ValueOf("captcha"),
		a.Ginger.GetController())
}

func (a *App[acc, prof, regReq, reg, f]) initiateAccount() {
	a.Account = account.New[acc, prof, f](
		a.Logger.WithTrace("account"),
		a.Registry.ValueOf("account"),
		a.Sql, a.Ginger.GetController(),
	)
}

func (a *App[acc, prof, regReq, reg, f]) initializePermission() {
	a.Permission = permission.Initialize(
		a.Logger.WithTrace("permission"),
		a.Sql,
	)
}

func (a *App[acc, prof, regReq, reg, f]) initiateSession() {
	a.Session = session.New(
		a.Logger.WithTrace("session"),
		a.Registry.ValueOf("session"),
		a.Cache,
	)
}

func (a *App[acc, prof, regReq, reg, f]) initiateLogin() {
	a.Login = login.New(
		a.Logger.WithTrace("login"),
		a.Registry.ValueOf("login"),
		a.Account.UseCase,
		a.Session.UseCase,
		a.Cache,
		a.Ginger.GetController(),
	)
}

func (a *App[acc, prof, regReq, reg, f]) initiateRegister() {
	a.Register = register.New[regReq, reg, acc](
		a.Logger.WithTrace("register"),
		a.Sql, a.Ginger.GetController(),
	)
}

func (a *App[acc, prof, regReq, reg, f]) initializeMonitoring() {
	a.Monitoring = monitoring.New(
		a.Logger.WithTrace("monitoring"),
		a.Ginger.GetController(),
	)
}
