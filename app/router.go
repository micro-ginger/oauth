package app

func (a *App[acc, reg]) registerRoutes() {
	rg := a.Ginger.NewRouterGroup("/")
	//
	// login
	loginGroup := rg.Group("/login")
	loginGroup.Create("", a.Login.Handler)
	//
	// accounts
	accountsGroup := rg.Group("/accounts")
	accountItemGroup := accountsGroup.Group("/:account_id")
	accountItemGroup.Read("",
		a.Authenticator.MustAuthenticate(),
		a.Account.GetHandler,
	)
	// account
	accountGroup := rg.Group("/account")
	accountGroup.Read("",
		a.Authenticator.MustAuthenticate(),
		a.Account.GetHandler,
	)
	//
}
