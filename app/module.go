package app

import (
	"github.com/micro-ginger/oauth/account"
	"github.com/micro-ginger/oauth/login"
)

func (a *App[acc]) initializeModules() {
	a.initiateAccount()
	a.initiateLogin()
}

func (a *App[acc]) initiateAccount() {
	a.Account = account.New[acc](
		a.Logger.WithTrace("account"),
		a.Sql, a.Ginger.GetController(),
	)
}

func (a *App[acc]) initiateLogin() {
	a.Login = login.New(
		a.Logger.WithTrace("login"),
		a.Registry.ValueOf("login"),
		a.Cache,
		a.Ginger.GetController(),
	)
}
