package app

func (a *app) registerRoutes() {
	_ = a.ginger.NewRouterGroup("/")
}
