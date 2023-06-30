package app

import "github.com/micro-ginger/oauth/account"

func (a *app[acc]) initializeModules() {
	a.initiateAccount()
}

func (a *app[acc]) initiateAccount() {
	a.Account = account.New(
		a.Logger.WithTrace("account"),
		a.Sql, a.Ginger.GetController(),
	)
}
