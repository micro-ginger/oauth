package app

import (
	"github.com/micro-ginger/oauth/account"
	"github.com/micro-ginger/oauth/login"
	"github.com/micro-ginger/oauth/register"
	"github.com/micro-ginger/oauth/session"
)

func (a *App[acc, regReq, reg]) initializeModules() {
	a.initiateAccount()
	a.initiateSession()
	a.initiateLogin()
	a.initiateRegister()
	//
	a.Register.Initialize(a.Account.UseCase)
}

func (a *App[acc, regReq, reg]) initiateAccount() {
	a.Account = account.New[acc](
		a.Logger.WithTrace("account"),
		a.Registry.ValueOf("account"),
		a.Sql, a.Ginger.GetController(),
	)
}

func (a *App[acc, regReq, reg]) initiateSession() {
	a.Session = session.New(
		a.Logger.WithTrace("session"),
		a.Registry.ValueOf("session"),
		a.Cache,
	)
}

func (a *App[acc, regReq, reg]) initiateLogin() {
	a.Login = login.New(
		a.Logger.WithTrace("login"),
		a.Registry.ValueOf("login"),
		a.Account.UseCase,
		a.Session.UseCase,
		a.Cache,
		a.Ginger.GetController(),
	)
}

func (a *App[acc, regReq, reg]) initiateRegister() {
	a.Register = register.New[regReq, reg, acc](
		a.Logger.WithTrace("register"),
		a.Sql, a.Ginger.GetController(),
	)
}
