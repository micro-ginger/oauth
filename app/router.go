package app

import "github.com/micro-ginger/oauth/global"

func (a *App[acc, prof, regReq, reg, f]) registerRoutes() {
	rg := a.Ginger.NewRouterGroup("/")
	//
	// chaptcha
	if a.Captcha != nil {
		rg.Create("/captcha/generate", a.Captcha.GenerateHandler)
	}
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
	accountGroup.Update("",
		a.Authenticator.MustHaveScope(global.ScopeUpdateProfile),
		a.Account.UpdateHandler,
	)
	// profile
	profileGroup := accountGroup.Group("/profile")
	profileGroup.Read("",
		a.Authenticator.MustHaveScope(global.ScopeReadProfile),
		a.Account.Profile.GetHandler,
	)
	profileGroup.Update("",
		a.Authenticator.MustHaveScope(global.ScopeUpdateProfile),
		a.Account.Profile.UpdateHandler,
	)
	profileGroup.Create("/photo",
		a.Authenticator.MustHaveScope(global.ScopeUpdateProfile),
		a.Account.Profile.PhotoUpdateHandler,
	)
	//
	// register
	registerGroup := rg.Group("/register")
	registerGroup.Create("",
		a.Authenticator.MustHaveScope(global.ScopeRegister),
		a.Register.RegisterHandler,
	)
	//
	// internal
	internalGroup := rg.Group("/internal")
	monitoringGroup := internalGroup.Group("/monitoring")
	monitoringGroup.Read("/health",
		a.Monitoring.Health,
	)
}
