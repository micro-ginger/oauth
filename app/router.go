package app

func (a *app[acc]) registerRoutes() {
	rg := a.Ginger.NewRouterGroup("/")
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
