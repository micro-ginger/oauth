package app

import (
	"github.com/micro-ginger/oauth/account"
	"github.com/micro-ginger/oauth/login"
	"github.com/micro-ginger/oauth/register"
)

func (a *App[acc, reg]) initializeModules() {
	a.initiateAccount()
	a.initiateLogin()
	a.initiateRegister()
}

func (a *App[acc, reg]) initiateAccount() {
	a.Account = account.New[acc](
		a.Logger.WithTrace("account"),
		a.Registry.ValueOf("account"),
		a.Sql, a.Ginger.GetController(),
	)
}

func (a *App[acc, reg]) initiateLogin() {
	a.Login = login.New(
		a.Logger.WithTrace("login"),
		a.Registry.ValueOf("login"),
		a.Cache,
		a.Ginger.GetController(),
	)
}

func (a *App[acc, reg]) initiateRegister() {
	a.Register = register.New[reg, acc](
		a.Logger.WithTrace("register"),
		a.Sql, a.Ginger.GetController(),
	)
}
