package app

import "github.com/micro-ginger/oauth/global"

func (a *App[acc, prof, regReq, reg]) registerRoutes() {
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
		a.Authenticator.MustHaveScope(global.ScopeReadAccount),
		a.Account.GetHandler,
	)
	// profile
	profileGroup := accountGroup.Group("/profile")
	profileGroup.Read("",
		a.Authenticator.MustHaveScope(global.ScopeReadProfile),
		a.Account.Profile.GetHandler,
	)
	//
	// register
	registerGroup := rg.Group("/register")
	registerGroup.Create("",
		a.Authenticator.MustHaveScope(global.ScopeRegister),
		a.Register.RegisterHandler,
	)
	//
}
