package app

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger"
	"github.com/micro-blonde/auth/authorization"
)

func (a *App[acc, reg]) initializeServer() {
	a.initializeGinger()
	a.initializeAuthenticator()
}

func (a *App[acc, reg]) initializeGinger() {
	logger := a.Logger.WithTrace("ginger")
	a.Ginger = ginger.NewServer(logger, a.Registry.ValueOf("gateway.http"))

	responder := a.Ginger.NewResponder()
	controller := gateway.NewController(responder).WithLanguageBundle(a.Language)
	a.Ginger.SetController(controller)
}

func (a *App[acc, reg]) initializeAuthenticator() {
	a.Authenticator = authorization.New[acc](
		a.Ginger, a.Registry.ValueOf("gateway.authorization"))
}
